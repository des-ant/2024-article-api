package main

import (
	"context"
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/des-ant/2024-article-api/internal/data"

	// Import the pq driver so that it can register itself with the database/sql
	// package. Alias this import to the blank identifier to stop the Go compiler
	// from complaining that it isn't being used.
	_ "github.com/lib/pq"
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
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

// Define an application struct to hold the dependencies for our HTTP handlers,
// helpers, and middleware.
type application struct {
	config config
	logger *slog.Logger
	daos   *data.DAOs
	wg     sync.WaitGroup
}

// parseFlags parses the command-line flags and returns a config struct.
func parseFlags(cfg *config) {
	// Define the flags for the HTTP server.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// Database	connection settings.
	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "PostgreSQL DSN")

	// Connection pool settings.
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")

	flag.Parse()
}

func main() {
	var cfg config

	parseFlags(&cfg)

	// Initialize a new structured logger which writes log entries to the standard
	// out stream.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create the database connection pool.
	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection pool established")

	// Declare an instance of the application struct, containing the config struct
	// and the logger.
	app := &application{
		config: cfg,
		logger: logger,
		daos:   data.NewDAOs(),
	}

	// Start the HTTP server.
	err = app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

// openDB opens a connection to the PostgreSQL database.
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
