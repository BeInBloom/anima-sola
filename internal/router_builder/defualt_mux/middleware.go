package defualtmux

import (
	"net/http"
	"slices"
)

func ContentType(types ...string) middleware {
	return func(next http.Handler) http.Handler {
		f := func(w http.ResponseWriter, r *http.Request) {
			contenType := r.Header.Get("Content-Type")

			if slices.Contains(types, contenType) {
				next.ServeHTTP(w, r)
				return
			}

			http.Error(w, "wrong contern type", http.StatusBadRequest)
		}

		return http.HandlerFunc(f)
	}
}
