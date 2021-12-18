package main

import (
	"context"
	"os"
	"time"

	"github.com/ademuanthony/dfctipper/app"
	"github.com/ademuanthony/dfctipper/postgres"
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

	client, err := getClient(&creds)
	if err != nil {
		log.Error("Error getting Twitter Client")
		log.Error(err)
		return
	}

	db, err := postgres.NewPgDb(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, os.Getenv("DEBUG_SQL") == "1")
	if err != nil {
		log.Errorf("pqsl: %v", err)
		return
	}

	ctx := context.Background()

	app.Start(ctx, db, client, b)

	log.Info("Bot starting...")

	b.Start()
}
