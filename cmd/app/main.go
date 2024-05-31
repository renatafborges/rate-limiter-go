package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renatafborges/rate-limiter-go/configs"
	"github.com/renatafborges/rate-limiter-go/internal/infra/cache"
	"github.com/renatafborges/rate-limiter-go/internal/infra/web"
)

const port = ":8080"

func init() {
	configs.LoadEnvConfig()
	configs.LoadRateLimitConfig()
	cache.LoadEnvCache()
}

func main() {

	router := mux.NewRouter()
	router.Use(web.RateLimiter)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to our system...")
	})

	http.Handle("/", router)
	log.Println("Starting server on: " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
