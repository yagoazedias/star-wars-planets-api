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

func TestPlanetCreateAndDeleteHandler(t *testing.T) {

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
		t.Errorf("Create Handler returned wrong status code: got %v want %v",
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

func TestPlanetUpdateAndDeleteHandler(t *testing.T) {

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
		t.Errorf("Was not possible to parse payload content from creation %v", err)
	}

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Create Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	req, err = http.NewRequest("PUT", fmt.Sprintf("/planet/id/%s", planet.ID.Hex()), bytes.NewBuffer([]byte(`{"name": "Mars","weather": "Hot","terrain": "Low"}`)))

	recorder := httptest.NewRecorder()
	handler = http.HandlerFunc(p.Update(formatter))

	req = mux.SetURLVars(req, map[string]string{"id": planet.ID.Hex()})

	handler.ServeHTTP(recorder, req)

	err = json.NewDecoder(recorder.Body).Decode(&planet)

	if err != nil {
		t.Errorf("Was not possible to parse payload content from updating %v", err)
	}

	if planet.Name != "Mars" {
		t.Errorf("Planet name was not updated: Got %v, expeted %v", planet.Name, "Mars")
	}

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Was not possible to update planet after create it. Got: %v, expeted: %v",
			status, http.StatusOK)
	}

	req, err = http.NewRequest("DELETE", fmt.Sprintf("/planet/id/%s", planet.ID.Hex()), nil)

	deleteRecorder := httptest.NewRecorder()
	handler = http.HandlerFunc(p.Delete(formatter))

	req = mux.SetURLVars(req, map[string]string{"id": planet.ID.Hex()})

	handler.ServeHTTP(deleteRecorder, req)

	if status := deleteRecorder.Code; status != http.StatusNoContent {
		t.Errorf("Was not possible to delete planet after create it. Got: %v, expeted: %v",
			status, http.StatusNoContent)
	}
}

func TestSearchHandler(t *testing.T) {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	p := Planet{}

	var planets []domain.Planet

	req, err := http.NewRequest("GET", "/planet", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(p.Search(formatter))

	handler.ServeHTTP(rr, req)

	err = json.NewDecoder(rr.Body).Decode(&planets)

	if err != nil {
		t.Errorf("Was not possible to parse payload content %v", err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Create Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPlanetLookupHandler(t *testing.T) {

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
		t.Errorf("Create Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	req, err = http.NewRequest("GET", fmt.Sprintf("/planet/id/%s", planet.ID.Hex()), nil)

	recorder := httptest.NewRecorder()
	handler = http.HandlerFunc(p.Lookup(formatter))

	req = mux.SetURLVars(req, map[string]string{"id": planet.ID.Hex()})

	handler.ServeHTTP(recorder, req)

	err = json.NewDecoder(recorder.Body).Decode(&planet)

	if err != nil {
		t.Errorf("Was not possible to parse payload content %v", err)
	}

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Was not possible to lookup planet after create it. Got: %v, expeted: %v",
			status, http.StatusOK)
	}

	req, err = http.NewRequest("DELETE", fmt.Sprintf("/planet/id/%s", planet.ID.Hex()), nil)

	deleteRecorder := httptest.NewRecorder()
	handler = http.HandlerFunc(p.Delete(formatter))

	req = mux.SetURLVars(req, map[string]string{"id": planet.ID.Hex()})

	handler.ServeHTTP(deleteRecorder, req)

	if status := deleteRecorder.Code; status != http.StatusNoContent {
		t.Errorf("Was not possible to delete planet after create it. Got: %v, expeted: %v",
			status, http.StatusNoContent)
	}
}