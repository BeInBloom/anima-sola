package di

import (
	"log/slog"

	"github.com/BeInBloom/anima-sol/internal/app"
	"github.com/BeInBloom/anima-sol/internal/handlers"
	"github.com/BeInBloom/anima-sol/internal/logger"
	maprepository "github.com/BeInBloom/anima-sol/internal/repository/map_repository"
	defaultmuxbuilder "github.com/BeInBloom/anima-sol/internal/router_builder/default_mux_builder"
	"github.com/BeInBloom/anima-sol/internal/router_builder/middlewares"
	"github.com/BeInBloom/anima-sol/models"
)

type di struct {
	cfg           *models.Config
	server        *app.ServerApp
	routerBuilder *defaultmuxbuilder.Builder
	handlers      *handlers.Handlers
	repo          *maprepository.Repo
	log           *slog.Logger
	middleware    *middlewares.Middleware
}

func New(cfg models.Config) di {
	return di{
		cfg: &cfg,
	}
}

func (d *di) Server() *app.ServerApp {
	if d.server == nil {
		d.server = app.New(models.ServerDeps{
			Host:       d.cfg.ServerConfig.Host,
			Port:       d.cfg.ServerConfig.Port,
			MuxBuilder: d.Builder(),
		})
	}

	return d.server
}

func (d *di) Builder() *defaultmuxbuilder.Builder {
	if d.routerBuilder == nil {
		d.routerBuilder = defaultmuxbuilder.New(
			models.MuxBuilderDeps{
				MwFactory:       d.Middlewares(),
				HandlersFactory: d.Handlers(),
			},
		)
	}

	return d.routerBuilder
}

func (d *di) Handlers() *handlers.Handlers {
	if d.handlers == nil {
		d.handlers = handlers.New(models.HandlerDeps{
			Repo: d.Repo(),
		})
	}

	return d.handlers
}

func (d *di) Repo() *maprepository.Repo {
	if d.repo == nil {
		d.repo = maprepository.New()
	}

	return d.repo
}

func (d *di) Logger() *slog.Logger {
	if d.log == nil {
		d.log = logger.New("dev")
	}

	return d.log
}

func (d *di) Middlewares() *middlewares.Middleware {
	if d.middleware == nil {
		d.middleware = middlewares.New(
			models.MiddlewareDeps{
				Log: d.Logger(),
			},
		)
	}

	return d.middleware
}
