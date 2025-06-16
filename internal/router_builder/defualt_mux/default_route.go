package defualtmux

import (
	"net/http"
	"regexp"
)

type (
	route struct {
		pattern *regexp.Regexp
		handler http.Handler
	}
)

func (r *route) With(mws ...middleware) {
	for _, mw := range mws {
		mw.Wrap(r.handler)
	}
}
