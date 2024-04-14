package http

import (
	"fmt"
	_struct "go-server/pkg/struct"
	"go-server/pkg/utils/logger"
	"go-server/pkg/utils/response"
	"net/http"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init(app *_struct.Application) (*http.Server, error) {
	router := chi.NewRouter()

	// Apply middleware to the router here
	if app.Config.Environment == "production" {
		router.Use(logger.Middleware())
	} else {
		router.Use(middleware.Logger)
		router.Use(middleware.RequestID)
		router.Use(middleware.RealIP)
		router.Use(corsMiddleware)
		router.Use(middleware.Recoverer)
	}

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	err := MapRoutes(router)
	if err != nil {
		return nil, err
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.HTTPPort),
		Handler: router,
	}
	return httpServer, nil
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		next.ServeHTTP(w, r)
	})
}

func MapRoutes(router *chi.Mux) error {
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]string{"message": "OK"})
	})

	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		// This will cause a panic
		var a []int
		fmt.Println(a[1])
	})

	// v1 := router.engine.Group("/api/v1")
	
	//userHttp.Routes(v1, s.database, s.validator)
	//productHttp.Routes(v1, s.database, s.validator, s.cache)
	//orderHttp.Routes(v1, s.database, s.validator)
	return nil
}
