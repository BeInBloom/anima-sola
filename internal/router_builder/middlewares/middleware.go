package middlewares

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"slices"

	"github.com/BeInBloom/anima-sol/models"
)

type middleware = func(http.Handler) http.Handler

type Middleware struct {
	log *slog.Logger
}

func New(deps models.MiddlewareDeps) *Middleware {
	log := deps.Log.With(
		slog.String("type", "Middleware"),
	)

	return &Middleware{
		log: log,
	}
}

func (m *Middleware) ContentType(types ...string) middleware {
	return func(next http.Handler) http.Handler {
		f := func(w http.ResponseWriter, r *http.Request) {
			contenType := r.Header.Get("Content-Type")

			if slices.Contains(types, contenType) {
				next.ServeHTTP(w, r)
				return
			}

			http.Error(w, "wrong contern type", http.StatusBadRequest)
		}

		return http.HandlerFunc(f)
	}
}

func (m *Middleware) Logger() middleware {
	log := m.log.With("middleware", "middleware_logger")

	return func(next http.Handler) http.Handler {
		f := func(w http.ResponseWriter, r *http.Request) {
			var bodyBuf bytes.Buffer
			body, err := io.ReadAll(io.TeeReader(r.Body, &bodyBuf))
			if err != nil {
				log.Error("error during buffer reading", slog.String("error", err.Error()))
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}

			log.Info("base_info",
				slog.String("url", r.URL.String()),
				slog.String("body", string(body)),
			)

			r.Body = io.NopCloser(&bodyBuf)
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(f)
	}
}
