package main

import (
	"go-server/internal/database"
	"go-server/internal/server/http"
	"go-server/pkg/adapter/redis"
	"go-server/pkg/config"
	_struct "go-server/pkg/struct"
	"go-server/pkg/utils/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

func main() {
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    // Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		errorLog.Fatalf("Failed to load configuration: %v", err)
	}

    // Init logger
	logger.Init(cfg.Environment)
    
    // Connect to database
	db, err := initDBConnection(cfg.Dsn)
	if err != nil {
		errorLog.Fatalf("Cannot connect to database %s", err)
	}
	defer db.Close()

    // Init application
	app := &_struct.Application{
		Db:        db,
		Cache: redis.New(redis.Config{
			Address:  cfg.RedisURI,
			Password: cfg.RedisPassword,
			Database: cfg.RedisDB,
		}),
		Config: cfg,
	}

    infoLog.Printf("Trying to start server ...")
	infoLog.Printf("health check: http://%s:%v/health ", cfg.Hostname, cfg.HTTPPort)
	startServer(app)
}

func initDBConnection(dsn string) (*pgxpool.Pool, error) {
	db, err := database.New(dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func startServer(app *_struct.Application) {
	go func() {
		srv, err := http.Init(app)
		if err != nil {
			log.Fatalf("Failed to init HTTP server: %v", err)
		}
		log.Printf("Starting server on %v 🚀", app.Config.HTTPPort)
		err = srv.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
			return
		}
	}()
	select {}
}