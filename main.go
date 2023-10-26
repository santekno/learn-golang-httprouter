package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//go:embed resources
var resources embed.FS

func main() {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")

	router.GET("/", SampleGetHandler)
	router.POST("/", SamplePostHandler)
	router.GET("/product/:id", GetUsedParamsHandler)
	router.GET("/product/:id/items/:itemId", NamedParameterHandler)
	router.GET("/images/*image", CatchAllParameterHandler)
	router.ServeFiles("/files/*filepath", http.FS(directory))

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	server.ListenAndServe()
}
