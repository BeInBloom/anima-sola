package defualtmux

import (
	"net/http"
	"regexp"
)

type middleware = func(http.Handler) http.Handler

type Mux struct {
	mw    []middleware
	routs []*route
}

func New() *Mux {
	return &Mux{
		mw:    make([]middleware, 0),
		routs: make([]*route, 0),
	}
}

func (m *Mux) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	for _, route := range m.routs {
		if route.pattern.MatchString(url) {
			route.handler.ServeHTTP(res, req)
			return
		}
	}

	http.NotFound(res, req)
}

func (m *Mux) Whit(mw ...middleware) {
	m.mw = append(m.mw, mw...)
}

func (m *Mux) SetRoute(pattern string, handler http.Handler) *route {
	re := regexp.MustCompile(pattern)
	route := &route{re, m.chain(handler)}
	m.routs = append(m.routs, route)

	return route
}

func (m *Mux) chain(h http.Handler) http.Handler {
	if len(m.mw) == 0 {
		return h
	}

	var handler http.Handler
	for _, mw := range m.mw {
		handler = mw(h)
	}

	return handler
}
