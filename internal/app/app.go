package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/BeInBloom/anima-sol/models"
)

var (
	ErrAppClosed     = errors.New("app was close")
	ErrOnClosingAggp = errors.New("closing error")
)

type (
	ServerApp struct {
		server *http.Server
	}

	routerBuilder interface {
		Router() http.Handler
	}
)

func New(deps models.ServerDeps) *ServerApp {
	s := http.Server{
		Addr:    fmt.Sprintf("%s:%d", deps.Host, deps.Port),
		Handler: deps.MuxBuilder.Router(),
	}

	return &ServerApp{
		server: &s,
	}
}

func (a *ServerApp) Run() error {
	const fn = "appr.Run"

	if err := a.server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return ErrAppClosed
		}

		return fmt.Errorf("%s:%w", fn, err)
	}

	return nil
}

func (a *ServerApp) Close() error {
	const fn = "app.Colose"

	ctx, cansel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cansel()

	if err := a.server.Shutdown(ctx); err != nil {
		return ErrOnClosingAggp
	}

	return nil
}
