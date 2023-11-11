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

func NamedParameterHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	text := "Product " + p.ByName("id") + " Item " + p.ByName("itemId")
	fmt.Fprint(w, text)
}

func CatchAllParameterHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	text := "Image " + p.ByName("image")
	fmt.Fprint(w, text)
}

func PanicHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("Oops panic")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "halaman tidak ditemukan")
}
