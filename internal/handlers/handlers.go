package handlers

import (
	"io"
	"math/rand/v2"
	"net/http"
	"net/url"
	"regexp"

	"github.com/BeInBloom/anima-sol/models"
)

type repo interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}

type Handlers struct {
	repo repo
}

func New(deps models.HandlerDeps) *Handlers {
	return &Handlers{
		repo: deps.Repo,
	}
}

func (h *Handlers) SetURLHandler() http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		url, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			h.handleError(w, []byte("cant read body"), http.StatusInternalServerError)
			return
		}

		if !isValidURL(string(url)) {
			h.handleError(w, []byte("not url"), http.StatusBadRequest)
			return
		}

		shortURL := getShortURL()

		if err := h.repo.Set(string(shortURL), string(url)); err != nil {
			h.handleError(w, []byte("cant save short url"), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(shortURL)
	}

	return http.HandlerFunc(f)
}

func (h *Handlers) GetURLHandler() http.Handler {
	re := regexp.MustCompile("^/([A-Za-z0-9]+)/?$")

	f := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		match := re.FindStringSubmatch(r.URL.Path)
		if len(match) < 2 {
			h.handleError(w, []byte("not correct query"), http.StatusBadRequest)
			return
		}

		shortURL := match[1]
		fullURL, err := h.repo.Get(shortURL)
		if err != nil {
			h.handleError(w, []byte("not found"), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fullURL))
	}

	return http.HandlerFunc(f)
}

func (h *Handlers) handleError(w http.ResponseWriter, message []byte, code int) {
	http.Error(w, string(message), code)
}

func getShortURL() []byte {
	const (
		length  = 8
		symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	)

	res := make([]byte, length)

	for i := range len(res) {
		res[i] = symbols[rand.IntN(len(symbols))]
	}

	return res
}

func isValidURL(uri string) bool {
	u, err := url.ParseRequestURI(uri)
	return err == nil && u.Scheme != "" && u.Host != ""
}
