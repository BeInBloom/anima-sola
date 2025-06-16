package models

import "net/http"

type (
	muxBuilder interface {
		Router() http.Handler
	}

	repo interface {
		Set(key string, value string) error
		Get(key string) (string, error)
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
