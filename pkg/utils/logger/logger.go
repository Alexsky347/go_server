package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	once sync.Once
)

func Init(env string) {
	once.Do(func() {
		if env == "production" {
			lumberjackLogger := &lumberjack.Logger{
				Filename:   "access.log",
				MaxSize:    5, // megabytes
				MaxBackups: 4,
				MaxAge:     14, //days
				Compress:   true,
			}

			multi := zerolog.MultiLevelWriter(lumberjackLogger, os.Stdout)
			log.Logger = zerolog.New(multi).With().Timestamp().Logger()
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		} else {
			log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Error().Msgf("panic recovered: %v", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()

			start := time.Now()
			path := r.URL.Path

			rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(rec, r)

			end := time.Now()
			latency := end.Sub(start)
			status := rec.status

			if status != http.StatusOK {
				log.Error().
					Str("path", path).
					Int("status", status).
					Dur("latency", latency).
					Str("method", r.Method).
					Str("clientIP", r.RemoteAddr).
					Msg("error")
			}
			//if status == http.StatusOK {
			//	log.Info().
			//		Str("path", path).
			//		Int("status", status).
			//		Dur("latency", latency).
			//		Str("method", r.Method).
			//		Str("clientIP", r.RemoteAddr).
			//		Msg("request")
			//}
		})
	}
}
