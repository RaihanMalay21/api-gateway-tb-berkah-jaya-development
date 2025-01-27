package main

import (
	"log"
	"net/http"

	"github.com/RaihanMalay21/api-gateway-tb-berkah-jaya-development/middlewares"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middlewares.CorsMiddlewares)
	r.PathPrefix("/customer").Handler(middlewares.ReverseProxy("http://localhost:8081"))
	r.PathPrefix("/access").Handler(middlewares.ReverseProxy("http://localhost:8082"))
	r.PathPrefix("/admin").Handler(middlewares.ReverseProxy("http://localhost:8083"))
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
