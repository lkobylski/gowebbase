package gowebbase

import (
	"log"
	"net/http"
)


func ExampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
		//TODO: implement yor middleware logic here

		next.ServeHTTP(w, req)
	})
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
		log.Printf("%s - %s\n", req.Method, req.URL)
		next.ServeHTTP(w, req)
	})
}