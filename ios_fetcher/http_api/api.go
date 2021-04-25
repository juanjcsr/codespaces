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

	HealthEndpoint HealthAPI
}

func NewApiClient(port string) *ApiClient {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	health := NewHealthAPI()

	return &ApiClient{
		router: r,
		port:   port,
		HealthEndpoint: *health,
	}
}

func (api *ApiClient) StartServer() {
	log.Println("starting server...")

	api.router.Get("/", api.HealthEndpoint.Root)

	http.ListenAndServe("0.0.0.0:"+api.port, api.router)
}
