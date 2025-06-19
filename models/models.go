package models

import (
	"log/slog"
	"net/http"
)

type (
	middleware = func(http.Handler) http.Handler
)

type (
	muxBuilder interface {
		Router() http.Handler
	}

	repo interface {
		Set(key string, value string) error
		Get(key string) (string, error)
	}

	middlewaresFactory interface {
		ContentType(types ...string) middleware
		Logger() middleware
	}

	handlersFactory interface {
		SetURLHandler() http.Handler
		GetURLHandler() http.Handler
	}
)

type (
	ServerDeps struct {
		Host       string
		Port       int
		MuxBuilder muxBuilder
	}

	HandlerDeps struct {
		Repo repo
	}

	MiddlewareDeps struct {
		Log *slog.Logger
	}

	MuxBuilderDeps struct {
		MwFactory       middlewaresFactory
		HandlersFactory handlersFactory
	}
)

type (
	Config struct {
		ServerConfig ServerConfig
	}

	ServerConfig struct {
		Host string
		Port int
	}
)
