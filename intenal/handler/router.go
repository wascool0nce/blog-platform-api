package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PostHandler interface {
	PostCreateHandler(w http.ResponseWriter, r *http.Request)
}

func NewRouter(postHandler PostHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Post("/posts/create", postHandler.PostCreateHandler)
	return r
}
