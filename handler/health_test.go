package handler

import (
	"encoding/json"
	"github.com/unrolled/render"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	Status string
}

func TestHealthCheckHandler(t *testing.T) {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	h := Health{}

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Check(formatter))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var r = Response{Status: "Not Ok"}
	expected := Response{Status: "Ok"}

	err = json.Unmarshal([]byte(rr.Body.String()), &r)

	if err != nil {
		log.Fatal(err)
	}

	if r != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}