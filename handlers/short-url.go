// handlers/handlers.go
package handlers

import (
	"encoding/json"
	"go-short-url/models"
	"go-short-url/services"
	"go-short-url/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UrlHandlerV1 struct {
	svc services.ShortenURLServicer
}

func NewShortenURLHandler(servicer services.ShortenURLServicer) chi.Router {
	// initializing the servicer
	v1 := &UrlHandlerV1{
		svc: servicer,
	}
	router := chi.NewRouter()
	router.Post("/", v1.ShortenURL)
	router.Get("/{shortURL}", v1.RedirectToURL)
	router.Get("/metrics", v1.Metrics)
	return router
}

func (v1 *UrlHandlerV1) ShortenURL(w http.ResponseWriter, r *http.Request) {
	data := RequestBody{}
	if err := render.Bind(r, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	domain, err := utils.GetDomain(data.URL)
	if err != nil {
		http.Error(w, "unable to process", http.StatusInternalServerError)
		return
	}

	shortURL, err := v1.svc.GetOrCreate(domain, data.URL)
	if err != nil {
		http.Error(w, "unable to process", http.StatusInternalServerError)
		return
	}

	shortURL = "http://localhost:8080/v1/urlshorter/" + shortURL
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(shortURL)); err != nil {
		http.Error(w, "unable to process", http.StatusInternalServerError)
		return
	}
}

func (v1 *UrlHandlerV1) RedirectToURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")

	originalURL, err := v1.svc.GetLongURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

func (v1 *UrlHandlerV1) Metrics(w http.ResponseWriter, r *http.Request) {
	byteData, err := json.Marshal(v1.svc.Metrics())
	if err != nil {
		http.Error(w, "unable to process", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(byteData); err != nil {
		http.Error(w, "unable to process", http.StatusInternalServerError)
		return
	}
}

type RequestBody struct {
	*models.RequestURL
}

func (u *RequestBody) Bind(r *http.Request) error {
	return u.Validate()
}
