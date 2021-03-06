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

	creationRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(p.Create(formatter))

	handler.ServeHTTP(creationRecorder, req)

	err = json.NewDecoder(creationRecorder.Body).Decode(&planet)

	if err != nil {
		t.Errorf("Was not possible to parse payload content %v", err)
	}

	if status := creationRecorder.Code; status != http.StatusCreated {
		t.Errorf("Create Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	req, err = http.NewRequest("POST", "/planet", bytes.NewBuffer([]byte(`{"name": "Earth","weather": "Hot","terrain": "Low"}`)))

	if err != nil {
		t.Fatal(err)
	}

	creationRecorder = httptest.NewRecorder()
	handler = http.HandlerFunc(p.Create(formatter))

	handler.ServeHTTP(creationRecorder, req)

	if status := creationRecorder.Code; status != http.StatusConflict {
		t.Errorf("Create Handler with not valid name returned wrong status code: got %v want %v",
			status, http.StatusConflict)
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

	creationRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(p.Create(formatter))

	handler.ServeHTTP(creationRecorder, req)

	err = json.NewDecoder(creationRecorder.Body).Decode(&planet)

	if err != nil {
		t.Errorf("Was not possible to parse payload content from creation %v", err)
	}

	if status := creationRecorder.Code; status != http.StatusCreated {
		t.Errorf("Create Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	req, err = http.NewRequest("PUT", fmt.Sprintf("/planet/id/%s", planet.ID.Hex()), bytes.NewBuffer([]byte(`{"name": "Mars","weather": "Hot","terrain": "Low"}`)))

	updateRecorder := httptest.NewRecorder()
	handler = http.HandlerFunc(p.Update(formatter))

	req = mux.SetURLVars(req, map[string]string{"id": planet.ID.Hex()})

	handler.ServeHTTP(updateRecorder, req)

	err = json.NewDecoder(updateRecorder.Body).Decode(&planet)

	if err != nil {
		t.Errorf("Was not possible to parse payload content from updating %v", err)
	}

	if planet.Name != "Mars" {
		t.Errorf("Planet name was not updated: Got %v, expeted %v", planet.Name, "Mars")
	}

	if status := updateRecorder.Code; status != http.StatusOK {
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

	searchRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(p.Search(formatter))

	handler.ServeHTTP(searchRecorder, req)

	err = json.NewDecoder(searchRecorder.Body).Decode(&planets)

	if err != nil {
		t.Errorf("Was not possible to parse payload content %v", err)
	}

	if status := searchRecorder.Code; status != http.StatusOK {
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

	creationRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(p.Create(formatter))

	handler.ServeHTTP(creationRecorder, req)

	err = json.NewDecoder(creationRecorder.Body).Decode(&planet)

	if err != nil {
		t.Errorf("Was not possible to parse payload content %v", err)
	}

	if status := creationRecorder.Code; status != http.StatusCreated {
		t.Errorf("Create Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	req, err = http.NewRequest("GET", fmt.Sprintf("/planet/id/%s", planet.ID.Hex()), nil)

	lookupRecorder := httptest.NewRecorder()
	handler = http.HandlerFunc(p.Lookup(formatter))

	req = mux.SetURLVars(req, map[string]string{"id": planet.ID.Hex()})

	handler.ServeHTTP(lookupRecorder, req)

	err = json.NewDecoder(lookupRecorder.Body).Decode(&planet)

	if err != nil {
		t.Errorf("Was not possible to parse payload content %v", err)
	}

	if status := lookupRecorder.Code; status != http.StatusOK {
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

func TestPlanetCreateForNotValidBody(t *testing.T) {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	buffers := []*bytes.Buffer{
		bytes.NewBuffer([]byte(`{"name": "Earth","weather": "Hot"}`)),
		bytes.NewBuffer([]byte(`{"name": "Earth","terrain": "Low"}`)),
		bytes.NewBuffer([]byte(`{"terrain": "Low"}`)),
		bytes.NewBuffer([]byte(`{"weather": "Hot","terrain": "Low"}`)),
		bytes.NewBuffer([]byte(`"Not Valid Json Object"`)),
	}

	for _, buffer := range buffers {
		p := Planet{}

		var planet domain.Planet

		req, err := http.NewRequest("POST", "/planet", buffer)

		if err != nil {
			t.Fatal(err)
		}

		creationRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(p.Create(formatter))

		handler.ServeHTTP(creationRecorder, req)

		err = json.NewDecoder(creationRecorder.Body).Decode(&planet)

		if err != nil {
			t.Errorf("Was not possible to parse payload content %v", err)
		}

		if status := creationRecorder.Code; status != http.StatusBadRequest {
			t.Errorf("Create Handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	}
}

func TestPlanetSearchByName(t *testing.T) {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	p := Planet{}

	var planet domain.Planet

	req, err := http.NewRequest("POST", "/planet", bytes.NewBuffer([]byte(`{"name": "Earth","weather": "Hot","terrain": "Low"}`)))

	if err != nil {
		t.Fatal(err)
	}

	creationRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(p.Create(formatter))

	handler.ServeHTTP(creationRecorder, req)

	err = json.NewDecoder(creationRecorder.Body).Decode(&planet)

	if err != nil {
		t.Errorf("Was not possible to parse payload content %v", err)
	}

	if status := creationRecorder.Code; status != http.StatusCreated {
		t.Errorf("Create Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	req, err = http.NewRequest("GET", "/planet?name=Earth", nil)

	if err != nil {
		t.Fatal(err)
	}

	searchRecorder := httptest.NewRecorder()
	handler = http.HandlerFunc(p.Search(formatter))

	handler.ServeHTTP(searchRecorder, req)

	var planets []domain.Planet

	err = json.NewDecoder(searchRecorder.Body).Decode(&planets)

	if err != nil {
		t.Errorf("Was not possible to parse payload content %v", err)
	}

	if len(planets) == 0 {
		t.Error("Handler haven't return any planet")
	}

	if planets[0].Name != planet.Name {
		t.Errorf("Planets names does not match: got %v want %v",
			planets[0].Name, planet.Name)
	}

	if status := searchRecorder.Code; status != http.StatusOK {
		t.Errorf("Create Handler returned wrong status code: got %v want %v",
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