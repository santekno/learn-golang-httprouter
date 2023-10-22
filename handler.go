package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SampleGetHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Hello Get")
}

func SamplePostHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Hello Post")
}

func GetUsedParamsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	text := "Product " + p.ByName("id")
	fmt.Fprint(w, text)
}
