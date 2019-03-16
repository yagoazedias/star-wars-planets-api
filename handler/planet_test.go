package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/yagoazedias/star-wars-planets-api/domain"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlanetCreateHandler(t *testing.T) {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	p := Planet{}

	var planet domain.Planet

	req, err := http.NewRequest("POST", "/planet", bytes.NewBuffer([]byte(`{"name": "Earth","weather": "Hot","terrain": "Low"}`)))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(p.Create(formatter))

	handler.ServeHTTP(rr, req)

	err = json.NewDecoder(rr.Body).Decode(&planet)

	if err != nil {
		t.Errorf("Was not possible to parse payload content %v", err)
	}

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	req, err = http.NewRequest("DELETE", fmt.Sprintf("/planet/id/%s", planet.ID.Hex()), nil)

	recorder := httptest.NewRecorder()
	handler = http.HandlerFunc(p.Delete(formatter))

	req = mux.SetURLVars(req, map[string]string{"id": planet.ID.Hex()})

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusNoContent {
		t.Errorf("Was not possible to delete planet after create it. Got: %v, expeted: %v",
			status, http.StatusNoContent)
	}
}
