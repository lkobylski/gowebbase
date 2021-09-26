package gowebbase

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexPageHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	srv := Server{}
	srv.Init()

	req,_ := http.NewRequest("GET", "/", nil)
	srv.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("index page return wrong request code %d but should return %d", rec.Code, http.StatusOK)
	}
}

func TestAboutPageHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	srv := Server{}
	srv.Init()

	req,_ := http.NewRequest("GET", "/about", nil)
	srv.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("index page return wrong request code %d but should return %d", rec.Code, http.StatusOK)
	}
}
