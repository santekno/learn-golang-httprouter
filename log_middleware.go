package main

import (
	"fmt"
	"net/http"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Before Execute Handler %s %s\n", r.Method, r.URL)
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After Execute Handler")
}
