package main

import (
	"go-short-url/configs"
	"go-short-url/services"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var router chi.Router

func TestMain(m *testing.M) {
	log.Println("Test main in progress")
	configs.ReadConfig()

	// initializing services
	urlSvc := services.NewShortenURLServicer()

	router = createRouter(urlSvc)

	rc := m.Run()
	os.Exit(rc)
}

func TestShouldShortenURL(t *testing.T) {
	req, _ := http.NewRequest("POST", "/v1/urlshorter", strings.NewReader(`{"url": "http://example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotEmpty(t, rr.Body.String())
}

func TestShouldNotShortenURL(t *testing.T) {
	req, _ := http.NewRequest("POST", "/v1/urlshorter", strings.NewReader(`{"url": ""}`))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.NotEmpty(t, rr.Body.String())
}

func TestRedirectURL(t *testing.T) {
	req, _ := http.NewRequest("POST", "/v1/urlshorter", strings.NewReader(`{"url": "http://example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	urls := strings.Split(rr.Body.String(), "/")

	req, _ = http.NewRequest("GET", "/v1/urlshorter/"+urls[len(urls)-1], nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMovedPermanently, rr.Code)
	assert.Equal(t, "http://example.com", rr.Header().Get("Location"))
}

func TestTopDomains(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/urlshorter/metrics", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotContains(t, rr.Body.String(), "test.com")

	req, _ = http.NewRequest("POST", "/v1/urlshorter", strings.NewReader(`{"url": "http://example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	req, _ = http.NewRequest("POST", "/v1/urlshorter", strings.NewReader(`{"url": "http://test.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	req, _ = http.NewRequest("GET", "/v1/urlshorter/metrics", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "example.com")
	assert.Contains(t, rr.Body.String(), "test.com")
}
