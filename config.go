package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/caarlos0/env"
	flags "github.com/jessevdk/go-flags"
)

const (
	defaultConfigFilename     = "dfctipper.conf"
	sampleConfigFileName      = "./sample-dfctipper.conf"
	defaultLogFilename        = "coinzion.log"
	defaultDataDirname        = "data"
	defaultLogLevel           = "info"
	defaultLogDirname         = "logs"
	defaultDbHost             = "0.0.0.0"
	defaultDbPort             = "5432"
	defaultDbUser             = "postgres"
	defaultDbPass             = "postgres"
	defaultDbName             = "coinzion"
	defaultAPIProto           = "http"
	defaultHost               = "0.0.0.0"
	defaultPort               = "7070"
	defaultServerHeader       = "dfctipper"
	defaultCacheControlMaxAge = 86400
)

var (
	defaultHomeDir    = "./"
	defaultConfigFile = filepath.Join(defaultHomeDir, defaultConfigFilename)
	defaultLogDir     = filepath.Join(defaultHomeDir, defaultLogDirname)
	defaultDataDir    = filepath.Join(defaultHomeDir, defaultDataDirname)
	dcrdHomeDir       = "./"
	defaultMaxLogZips = 16
)

type config struct {
	// General application behavior
	HomeDir     string `short:"A" long:"appdata" description:"Path to application home directory" env:"PDANALYTICS_APPDATA_DIR"`
	ConfigFile  string `short:"C" long:"configfile" description:"Path to configuration file" env:"PDANALYTICS_CONFIG_FILE"`
	DataDir     string `short:"b" long:"datadir" description:"Directory to store data" env:"PDANALYTICS_DATA_DIR"`
	LogDir      string `long:"logdir" description:"Directory to log output." env:"PDANALYTICS_LOG_DIR"`
	MaxLogZips  int    `long:"max-log-zips" description:"The number of zipped log files created by the log rotator to be retained. Setting to 0 will keep all."`
	OutFolder   string `short:"f" long:"outfolder" description:"Folder for file outputs" env:"PDANALYTICS_OUT_FOLDER"`
	ShowVersion bool   `short:"V" long:"version" description:"Display version information and exit"`
	DebugLevel  string `short:"d" long:"debuglevel" description:"Logging level {trace, debug, info, warn, error, critical}" env:"PDANALYTICS_LOG_LEVEL"`
	Quiet       bool   `short:"q" long:"quiet" description:"Easy way to set debuglevel to error" env:"PDANALYTICS_QUIET"`

	// Postgresql Configuration
	DBHost string `long:"dbhost" description:"Database host" env:"DBHOST"`
	DBPort string `long:"dbport" description:"Database port" env:"DBPORT"`
	DBUser string `long:"dbuser" description:"Database username" env:"DBUSER"`
	DBPass string `long:"dbpass" description:"Database password" env:"DBPASS"`
	DBName string `long:"dbname" description:"Database name" env:"DBNAME"`

	// Twitter credentials
	ConsumerKey       string `env:"CONSUMER_KEY"`
	ConsumerSecret    string `env:"CONSUMER_SECRET"`
	AccessToken       string `env:"ACCESS_TOKEN"`
	AccessTokenSecret string `env:"ACCESS_TOKEN_SECRET"`

	TelegramAuth string `env:"TELEGRAM_AUTH"`

	BSCNode          string `long:"MAINNET_NODE_ADDRESS" env:"MAINNET_NODE_ADDRESS"`
	MasterAddressKey string `long:"MASTER_ADDRESS_KEY" env:"MASTER_ADDRESS_KEY"`
	MasterAddress    string `long:"MASTER_ADDRESS" env:"MASTER_ADDRESS"`

	// EMAIL
	MailgunDomain string `long:"mailgudomain" env:"MAILGUNDOMAIN"`
	MailgunAPIKey string `long:"mailgunapikey" env:"MAILGUNAPIKEY"`

	EnableWeb     string `long:"enableweb" env:"ENABLEWEB"`
	EnableTwitter string `long:"enabletwitter" env:"ENABLETWITTER"`

	// API/server
	APIProto           string `long:"apiproto" description:"Protocol for API (http or https)" env:"PDANALYTICS_ENABLE_HTTPS"`
	APIListen          string `long:"apilisten" description:"Listen address for API. default localhost:7777, :17778 testnet, :17779 simnet" env:"PDANALYTICS_LISTEN_URL"`
	ServerHeader       string `long:"server-http-header" description:"Set the HTTP response header Server key value. Valid values are \"off\", \"version\", or a custom string."`
	CacheControlMaxAge int    `long:"cachecontrol-maxage" description:"Set CacheControl in the HTTP response header to a value in seconds for clients to cache the response. This applies only to FileServer routes." env:"DCRDATA_MAX_CACHE_AGE"`
}

func defaultConfig() config {
	cfg := config{
		HomeDir:    defaultHomeDir,
		DataDir:    defaultDataDir,
		LogDir:     defaultLogDir,
		DBHost:     defaultDbHost,
		DBPort:     defaultDbPort,
		DBUser:     defaultDbUser,
		DBPass:     defaultDbPass,
		DBName:     defaultDbName,
		MaxLogZips: defaultMaxLogZips,
		ConfigFile: defaultConfigFile,
		DebugLevel: defaultLogLevel,

		APIProto:           defaultAPIProto,
		CacheControlMaxAge: defaultCacheControlMaxAge,
		ServerHeader:       defaultServerHeader,
	}

	return cfg
}

// normalizeNetworkAddress checks for a valid local network address format and
// adds default host and port if not present. Invalidates addresses that include
// a protocol identifier.
func normalizeNetworkAddress(a, defaultHost, defaultPort string) (string, error) {
	if strings.Contains(a, "://") {
		return a, fmt.Errorf("Address %s contains a protocol identifier, which is not allowed", a)
	}
	if a == "" {
		return defaultHost + ":" + defaultPort, nil
	}
	host, port, err := net.SplitHostPort(a)
	if err != nil {
		if strings.Contains(err.Error(), "missing port in address") {
			normalized := a + ":" + defaultPort
			host, port, err = net.SplitHostPort(normalized)
			if err != nil {
				return a, fmt.Errorf("Unable to address %s after port resolution: %v", normalized, err)
			}
		} else {
			return a, fmt.Errorf("Unable to normalize address %s: %v", a, err)
		}
	}
	if host == "" {
		host = defaultHost
	}
	if port == "" {
		port = defaultPort
	}
	return host + ":" + port, nil
}

// loadConfig initializes and parses the config using a config file and command
// line options.
func loadConfig() (*config, error) {
	loadConfigError := func(err error) (*config, error) {
		return nil, err
	}

	// Default config
	cfg := defaultConfig()
	defaultConfigNow := defaultConfig()

	// Load settings from environment variables.
	err := env.Parse(&cfg)
	if err != nil {
		return loadConfigError(err)
	}

	// If appdata was specified but not the config file, change the config file
	// path, and record this as the new default config file location.
	if defaultHomeDir != cfg.HomeDir && defaultConfigNow.ConfigFile == cfg.ConfigFile {
		cfg.ConfigFile = filepath.Join(cfg.HomeDir, defaultConfigFilename)
		// Update the defaultConfig to avoid an error if the config file in this
		// "new default" location does not exist.
		defaultConfigNow.ConfigFile = cfg.ConfigFile
	}

	// Pre-parse the command line options to see if an alternative config file
	// or the version flag was specified. Override any environment variables
	// with parsed command line flags.
	preCfg := cfg
	preParser := flags.NewParser(&preCfg, flags.HelpFlag|flags.PassDoubleDash)
	_, flagerr := preParser.Parse()

	if flagerr != nil {
		e, ok := flagerr.(*flags.Error)
		if !ok || e.Type != flags.ErrHelp {
			preParser.WriteHelp(os.Stderr)
		}
		if ok && e.Type == flags.ErrHelp {
			preParser.WriteHelp(os.Stdout)
			os.Exit(0)
		}
		return loadConfigError(flagerr)
	}

	// Show the version and exit if the version flag was specified.
	appName := filepath.Base(os.Args[0])
	appName = strings.TrimSuffix(appName, filepath.Ext(appName))
	if preCfg.ShowVersion {
		fmt.Printf("%s version 1.0 (Go version %s)\n", appName, runtime.Version())
		os.Exit(0)
	}

	port := defaultPort
	log.Info("Env $PORT :", os.Getenv("PORT"))
	if os.Getenv("PORT") != "" {
		_, err = strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Critical(err)
			log.Critical("$PORT must be set")
		}
		port = os.Getenv("PORT")
	}
	// Check the supplied APIListen address
	if cfg.APIListen == "" {
		cfg.APIListen = defaultHost + ":" + port
	} else {
		cfg.APIListen, err = normalizeNetworkAddress(cfg.APIListen, defaultHost, port)
		if err != nil {
			return loadConfigError(err)
		}
	}

	return &cfg, nil
}
