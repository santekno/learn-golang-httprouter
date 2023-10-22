package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", SampleGetHandler)
	router.POST("/", SamplePostHandler)
	router.GET("/product/:id", GetUsedParamsHandler)

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	server.ListenAndServe()
}
