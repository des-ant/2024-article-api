package main

import (
	"flag"
	"log/slog"
	"os"
	"sync"

	"github.com/des-ant/2024-article-api/internal/data"
)

// Declare a string containing the application version number.
const version = "1.0.0"

// Define a config struct to hold all the configuration settings for our
// application.
// Currently includes:
// - Network port for the server
// - Operating environment (development, staging, production, etc.)
type config struct {
	port int
	env  string
}

// Define an application struct to hold the dependencies for our HTTP handlers,
// helpers, and middleware.
type application struct {
	config config
	logger *slog.Logger
	daos   *data.DAOs
	wg     sync.WaitGroup
}

// parseFlags reads the value of the port and env command-line flags into the config struct.
func parseFlags(cfg *config) {
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()
}

func main() {
	var cfg config

	parseFlags(&cfg)

	// Initialize a new structured logger which writes log entries to the standard
	// out stream.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Declare an instance of the application struct, containing the config struct
	// and the logger.
	app := &application{
		config: cfg,
		logger: logger,
		daos:   data.NewDAOs(),
	}

	// Start the HTTP server.
	err := app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
