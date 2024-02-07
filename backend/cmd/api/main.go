package main

import (
	"encoding/gob"
	"expvar"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
	"v0/internal/data"
	"v0/internal/jsonlog"
)

const version = "1.0.0"

type config struct {
	port  int
	debug bool
	db    struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	secret struct {
		HMC               string
		secretKey         []byte
		sessionExpiration time.Duration
	}
	tokenExpiration struct {
		durationString string
		duration       time.Duration
	}
}

type application struct {
	config config
	models data.Models
	logger *jsonlog.Logger
	wg     sync.WaitGroup
}

func main() {
	gob.Register(&data.UserID{})

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	cfg, err := updateConfigWithEnvVariables()
	if err != nil {
		logger.PrintFatal(err, nil, cfg.debug)
	}

	db, err := openDB(*cfg)

	if err != nil {
		logger.PrintFatal(err, nil, cfg.debug)
	}

	defer db.Close()

	logger.PrintInfo("database connection pool established", map[string]string{
		"debug": strconv.FormatBool(cfg.debug),
	}, cfg.debug)

	if err != nil {
		logger.PrintFatal(err, nil, cfg.debug)
	}

	expvar.NewString("version").Set(version)
	expvar.Publish("goroutines", expvar.Func(func() interface{} {
		return runtime.NumGoroutine()
	}))
	expvar.Publish("database", expvar.Func(func() interface{} {
		return db.Stats()
	}))
	expvar.Publish("timestamp", expvar.Func(func() interface{} {
		return time.Now().Unix()
	}))

	app := &application{
		config: *cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil, cfg.debug)
	}

}
