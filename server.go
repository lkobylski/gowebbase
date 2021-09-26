package gowebbase

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	// Server options
	options ServerOpts

	//DB connection instance
	db *gorm.DB

	templates map[string]*template.Template

	// Http Router
	router *mux.Router
}

type ServerOpts struct {
	Port int
	DSN string
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		options:   opts,
	}
}

func (s *Server) Init() {
	s.router = mux.NewRouter().StrictSlash(true)

	//Use log middleware
	s.router.Use(LogMiddleware)

	s.router.Handle("/", IndexPageHandler(s)).Methods("GET")
	s.router.PathPrefix("/public/").HandlerFunc(StaticFilesHandler)
	s.router.Handle("/about", AboutPageHandler(s)).Methods("GET")

	//Build templates
	if err := s.prepareTemplates(); err != nil {
		log.Panic(err)
	}

	//Add db connection

	//Setup tables
}

func (s Server) responseHtmlError(w http.ResponseWriter, status int, err error) {

	http.Error(w, err.Error(), status)
}

func (s *Server) responseHtmlTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := s.templates[name]
	if !ok {
		s.responseHtmlError(w, http.StatusInternalServerError, errors.New(fmt.Sprintf("template %s doesn't exists", name)))
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf=8")
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		s.responseHtmlError(w, http.StatusInternalServerError, err)
		return
	}
}

func (s *Server) Run() error {
	//close db connection
	//TODO add gracefull shoutdown
	log.Printf("server is running on :%d\n", s.options.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.options.Port), s.router))
	return nil
}

func (s *Server) prepareTemplates() (err error) {
	if s.templates == nil {
		s.templates = make(map[string]*template.Template)
	}

	baseLayout := "templates/layout.html"
	pages := []string{
		"templates/index.html",
		"templates/about.html",
	}

	for _, f := range pages {
		name := strings.ReplaceAll(f, "templates/", "")
		s.templates[name], err = template.New(name).ParseFiles(f, baseLayout)
		if err != nil {
			return err
		}
	}

	return nil
}
