package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/BeInBloom/anima-sol/models"
)

var (
	ErrAppClosed = errors.New("app was close")
)

type (
	app struct {
		server *http.Server
	}

	routerBuilder interface {
		Router() http.Handler
	}
)

func New(deps models.ServerDeps) app {
	u := &url.URL{
		Scheme: deps.Scheme,
		Host:   fmt.Sprintf("%s:%d", deps.Host, deps.Port),
	}

	s := http.Server{
		Addr:    u.Host,
		Handler: deps.MuxBuilder.Router(),
	}

	return app{
		server: &s,
	}
}

func (a *app) Run() error {
	return a.server.ListenAndServe()
}

func (a *app) Close() error {
	ctx, cansel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cansel()
	return a.server.Shutdown(ctx)
}
