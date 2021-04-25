package main

import "github.com/juanjcsr/ios_fetcher/http_api"

func main() {
	api := http_api.NewApiClient("3000")

	api.StartServer()
}
