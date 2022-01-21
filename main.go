package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ademuanthony/dfctipper/app"
	"github.com/ademuanthony/dfctipper/postgres"
	"github.com/ademuanthony/dfctipper/web"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
		return
	}

	// Parse the configuration file, and setup logger.
	cfg, err := loadConfig()
	if err != nil {
		log.Errorf("Failed to load pdanalytics config: %s\n", err.Error())
		return
	}

	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		// URL: "http://195.129.111.17:8012",
		URL: "https://api.telegram.org",

		Token:  cfg.TelegramAuth,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		// Poller: &tb.Webhook{
		// 	Listen:         "localhost:9000",
		// 	MaxConnections: 100,
		// 	Endpoint: &tb.WebhookEndpoint{
		// 		PublicURL: "https://0bfbd30c0d7b.ngrok.io",
		// 	},
		// },
	})

	if err != nil {
		log.Error(err)
		return
	}

	log.Info("Go-Twitter Bot v0.01")
	creds := Credentials{
		AccessToken:       cfg.AccessToken,
		AccessTokenSecret: cfg.AccessTokenSecret,
		ConsumerKey:       cfg.ConsumerKey,
		ConsumerSecret:    cfg.ConsumerSecret,
	}

	var client *twitter.Client
	if cfg.EnableTwitter == "1" {
		client, err = getClient(&creds)

		if err != nil {
			log.Error("Error getting Twitter Client")
			log.Error(err)
			return
		}
	}

	db, err := postgres.NewPgDb(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, os.Getenv("DEBUG_SQL") == "1")
	if err != nil {
		log.Errorf("pqsl: %v", err)
		return
	}

	ctx := context.Background()

	log.Info("Nade address", cfg.BSCNode)
	ethClient, err := ethclient.Dial(cfg.BSCNode)
	if err != nil {
		log.Error("ethClient", err)
		return
	}

	defer ethClient.Close()

	var wg sync.WaitGroup
	webMux := chi.NewRouter()
	webServer, err := web.NewServer(web.Config{
		CacheControlMaxAge: int64(cfg.CacheControlMaxAge),
		Viewsfolder:        "./views",
		AssetsFolder:       "./public",
		ReloadHTML:         true,
	}, webMux)
	if err != nil {
		log.Errorf("failed to create new web server (templates missing?) - %v", err)
		return
	}

	webServer.MountAssetPaths("/", "./public")

	if err := app.Start(ctx, webServer, db, client, ethClient, app.BlockchainConfig{
		BSCNode: cfg.BSCNode, MasterAddressKey: cfg.MasterAddressKey, MasterAddress: cfg.MasterAddress,
	}, b, cfg.MailgunDomain, cfg.MailgunAPIKey, cfg.EnableWeb == "1", cfg.EnableTwitter == "1"); err != nil {
		log.Error(err)
		return
	}

	// buildExplorer start the web server when its convenient for we are
	// starting here if the block explorer is disable.
	// The action here assumes that all other modules has being configured
	webServer.BuildRoute()
	listenAndServeProto(ctx, &wg, cfg.APIListen, cfg.APIProto, webMux)

	log.Info("Bot starting...")

	b.Start()
}

func listenAndServeProto(ctx context.Context, wg *sync.WaitGroup, listen, proto string, mux http.Handler) {
	// Try to bind web server
	server := http.Server{
		Addr:         listen,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,  // slow requests should not hold connections opened
		WriteTimeout: 60 * time.Second, // hung responses must die
	}

	// Add the graceful shutdown to the waitgroup.
	wg.Add(1)
	go func() {
		// Start graceful shutdown of web server on shutdown signal.
		<-ctx.Done()

		// We received an interrupt signal, shut down.
		log.Infof("Gracefully shutting down web server...")
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners.
			log.Infof("HTTP server Shutdown: %v", err)
		}

		// wg.Wait can proceed.
		wg.Done()
	}()

	log.Infof("Now serving the explorer and APIs on %s://%v/", proto, listen)
	// Start the server.
	go func() {
		var err error
		if proto == "https" {
			err = server.ListenAndServeTLS("pdanalytics.cert", "pdanalytics.key")
		} else {
			err = server.ListenAndServe()
		}
		// If the server dies for any reason other than ErrServerClosed (from
		// graceful server.Shutdown), log the error and request pdanalytics be
		// shutdown.
		if err != nil && err != http.ErrServerClosed {
			log.Errorf("Failed to start server: %v", err)
			requestShutdown()
		}
	}()

	// If the server successfully binds to a listening port, ListenAndServe*
	// will block until the server is shutdown. Wait here briefly so the startup
	// operations in main can have a chance to bail out.
	time.Sleep(250 * time.Millisecond)
}

// shutdownRequested checks if the Done channel of the given context has been
// closed. This could indicate cancellation, expiration, or deadline expiry. But
// when called for the context provided by withShutdownCancel, it indicates if
// shutdown has been requested (i.e. via requestShutdown).
func shutdownRequested(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// shutdownRequest is used to initiate shutdown from one of the
// subsystems using the same code paths as when an interrupt signal is received.
var shutdownRequest = make(chan struct{})

// shutdownSignal is closed whenever shutdown is invoked through an interrupt
// signal or from an JSON-RPC stop request.  Any contexts created using
// withShutdownChannel are cancelled when this is closed.
var shutdownSignal = make(chan struct{})

// signals defines the signals that are handled to do a clean shutdown.
// Conditional compilation is used to also include SIGTERM on Unix.
var signals = []os.Signal{os.Interrupt}

// withShutdownCancel creates a copy of a context that is cancelled whenever
// shutdown is invoked through an interrupt signal or from an JSON-RPC stop
// request.
func withShutdownCancel(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-shutdownSignal
		cancel()
	}()
	return ctx
}

// requestShutdown signals for starting the clean shutdown of the process
// through an internal component (such as through the JSON-RPC stop request).
func requestShutdown() {
	shutdownRequest <- struct{}{}
}

// shutdownListener listens for shutdown requests and cancels all contexts
// created from withShutdownCancel.  This function never returns and is intended
// to be spawned in a new goroutine.
func shutdownListener() {
	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, signals...)

	// Listen for the initial shutdown signal
	select {
	case sig := <-interruptChannel:
		log.Infof("Received signal (%s). Shutting down...", sig)
	case <-shutdownRequest:
		log.Info("Shutdown requested. Shutting down...")
	}

	// Cancel all contexts created from withShutdownCancel.
	close(shutdownSignal)

	// Listen for any more shutdown signals and log that shutdown has already
	// been signaled.
	for {
		select {
		case <-interruptChannel:
		case <-shutdownRequest:
		}
		log.Info("Shutdown signaled. Already shutting down...")
	}
}
