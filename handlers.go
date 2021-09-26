package gowebbase

import (
	"net/http"
)

func IndexPageHandler(s *Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		s.responseHtmlTemplate(w, "index.html", nil)
	})
}

func AboutPageHandler(s *Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		s.responseHtmlTemplate(w, "about.html", nil)
	})
}

func StaticFilesHandler(w http.ResponseWriter, req *http.Request) {
	u := req.URL.Path[1:]
	http.ServeFile(w, req, u)
}