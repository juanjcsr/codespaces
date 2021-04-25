package http_api

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ApiClient struct {
	router *chi.Mux
	port   string
}

func NewApiClient(port string) *ApiClient {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	return &ApiClient{
		router: r,
		port:   port,
	}
}

func (api *ApiClient) StartServer() {
	log.Println("starting server...")
	http.ListenAndServe("0.0.0.0:"+api.port, api.router)
}
