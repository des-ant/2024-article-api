package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// Declare a string containing the application version number.
// TODO: auto-generate at build time.
const version = "1.0.0"

// Define a config struct to hold all the configuration settings for our
// application.
// TODO: Add more settings as needed.
// Currently includes:
// - Network port for the server
// - Operating environment (development, staging, production, etc.)
// TODO: Read settings from command-line flags at startup.
type config struct {
	port int
	env  string
}

// Define an application struct to hold the dependencies for our HTTP handlers,
// helpers, and middleware.
type application struct {
	config config
	logger *slog.Logger
}

// parseFlags reads the value of the port and env command-line flags into the
// config struct.
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
	}

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", cfg.port),
		Handler:     app.routes(),
		IdleTimeout: time.Minute,
		// Set timeouts to prevent slow clients from consuming resources.
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// Start the HTTP server.
	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
