package defaultmuxbuilder

import (
	"net/http"

	defualtmux "github.com/BeInBloom/anima-sol/internal/router_builder/defualt_mux"
)

type (
	middleware interface {
		Wrap(http.Handler) http.Handler
	}

	mux interface {
		http.Handler
		SetRoute(pattern string, handler http.Handler)
		WhitMiddleware(mw ...middleware)
	}

	builder struct {
		mux http.Handler
	}
)

func New() *builder {
	mux := defualtmux.New()

	mux.SetRoute("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}))

	return &builder{
		mux: mux,
	}
}

func (b *builder) Router() http.Handler {
	return b.mux
}
