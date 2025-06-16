package defaultmuxbuilder

import (
	"net/http"

	defualtmux "github.com/BeInBloom/anima-sol/internal/router_builder/defualt_mux"
)

type Builder struct {
	mux http.Handler
}

type handlerFactory interface {
	SetURLHandler() http.Handler
	GetURLHandler() http.Handler
}

func New(hf handlerFactory) *Builder {
	mux := defualtmux.New()

	mux.SetRoute("^/$", hf.SetURLHandler())

	mux.SetRoute("^/([A-Za-z0-9]+)/?$", hf.GetURLHandler())

	return &Builder{
		mux: mux,
	}
}

func (b *Builder) Router() http.Handler {
	return b.mux
}
