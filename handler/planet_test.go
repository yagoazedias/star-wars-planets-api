package handler

import (
	"bytes"
	"github.com/unrolled/render"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlanetCreateHandler(t *testing.T) {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	p := Planet{}

	req, err := http.NewRequest("POST", "/planet", bytes.NewBuffer([]byte(`{"name": "Earth","weather": "Hot","terrain": "Low"}`)))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(p.Create(formatter))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
