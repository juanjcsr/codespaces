package http_api

import "net/http"

type HealthAPI struct {
	Root http.HandlerFunc
}

func NewHealthAPI() *HealthAPI {
	return &HealthAPI{
		Root: rootEndpoint,
	}
} 

func rootEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
