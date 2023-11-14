package main

import (
	"embed"
	"fmt"
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

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "Panic ", i)
	}

	router.GET("/panic", PanicHandler)
	// handler custom for not found
	router.NotFound = http.HandlerFunc(NotFoundHandler)
	// handler custom for method not allowdd
	router.MethodNotAllowed = http.HandlerFunc(MethodNotAllowedHandler)

	middleware := &LogMiddleware{router}

	server := http.Server{
		Handler: middleware,
		Addr:    "localhost:8080",
	}

	server.ListenAndServe()
}
