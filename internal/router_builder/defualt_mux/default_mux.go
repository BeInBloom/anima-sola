package defualtmux

import (
	"net/http"
	"regexp"
)

type (
	mux struct {
		mw    []middleware
		routs []*route
	}

	middleware interface {
		Wrap(http.Handler) http.Handler
	}
)

func New() *mux {
	return &mux{
		mw:    make([]middleware, 0),
		routs: make([]*route, 0),
	}
}

func (m *mux) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	for _, route := range m.routs {
		if route.pattern.MatchString(url) {
			route.handler.ServeHTTP(res, req)
			return
		}
	}

	http.NotFound(res, req)
}

func (m *mux) WhitMiddleware(mw ...middleware) {
	m.mw = append(m.mw, mw...)
}

func (m *mux) SetRoute(pattern string, handler http.Handler) *route {
	re := regexp.MustCompile(pattern)
	route := &route{re, m.chain(handler)}
	m.routs = append(m.routs, route)

	return route
}

func (m *mux) chain(h http.Handler) http.Handler {
	if len(m.mw) == 0 {
		return h
	}

	var handler http.Handler
	for _, mw := range m.mw {
		handler = mw.Wrap(h)
	}

	return handler
}
