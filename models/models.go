package models

import "net/http"

type (
	muxBuilder interface {
		Router() http.Handler
	}
)

type (
	ServerDeps struct {
		Scheme     string
		Host       string
		Port       int
		MuxBuilder muxBuilder
	}
)
