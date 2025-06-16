package di

import (
	"github.com/BeInBloom/anima-sol/internal/app"
	"github.com/BeInBloom/anima-sol/internal/handlers"
	maprepository "github.com/BeInBloom/anima-sol/internal/repository/map_repository"
	defaultmuxbuilder "github.com/BeInBloom/anima-sol/internal/router_builder/default_mux_builder"
	"github.com/BeInBloom/anima-sol/models"
)

type di struct {
	cfg           *models.Config
	server        *app.ServerApp
	routerBuilder *defaultmuxbuilder.Builder
	handlers      *handlers.Handlers
	repo          *maprepository.Repo
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
		d.routerBuilder = defaultmuxbuilder.New(d.Handlers())
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
