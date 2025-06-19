package defaultmuxbuilder

import (
	"net/http"

	defualtmux "github.com/BeInBloom/anima-sol/internal/router_builder/defualt_mux"
	"github.com/BeInBloom/anima-sol/models"
)

type Builder struct {
	mux http.Handler
}

func New(deps models.MuxBuilderDeps) *Builder {
	mux := defualtmux.New()

	mux.Whit(deps.MwFactory.Logger())

	mux.SetRoute("^/$", deps.HandlersFactory.SetURLHandler())
	mux.SetRoute("^/([A-Za-z0-9]+)/?$", deps.HandlersFactory.GetURLHandler())

	return &Builder{
		mux: mux,
	}
}

func (b *Builder) Router() http.Handler {
	return b.mux
}
