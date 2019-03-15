package service

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/yagoazedias/star-wars-planets-api/client"
	"github.com/yagoazedias/star-wars-planets-api/domain"
	"github.com/yagoazedias/star-wars-planets-api/helpers"
	"github.com/yagoazedias/star-wars-planets-api/repository"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

type Planet struct {
	Repository repository.Planet
}

func (m *Planet) Search(offset string, limit string) ([]domain.Planet, int, error)  {

	parsedOffset, err := strconv.Atoi(offset)
	parsedLimit, err := strconv.Atoi(limit)

	planets, err := m.Repository.Search(parsedOffset, parsedLimit)

	if err != nil {
		return nil, http.StatusInternalServerError, helpers.NewError("An unexpected error has occurred")
	}

	for i, planet := range planets {
		swapi := client.Swapi{}

		planetAttendanceInFilms, err := swapi.GetPlanetAttendance(&planet)

		if err != nil {
			fmt.Printf("Error at swapi integration %s", err)
		}

		planets[i].Count = planetAttendanceInFilms
	}

	return planets, http.StatusOK, nil
}

func (m *Planet) Lookup(request *http.Request) (*domain.Planet, int, error)  {

	vars := mux.Vars(request)

	if vars["id"] == "" {
		return nil, http.StatusBadRequest, helpers.NewError("Url Param 'id' is missing")
	}

	isValid := govalidator.IsMongoID(vars["id"])

	if !isValid {
		return nil, http.StatusBadRequest, helpers.NewError("Not a valid id")
	}

	planet, err := m.Repository.Lookup(bson.ObjectIdHex(vars["id"]))

	if err != nil {
		return nil, http.StatusNotFound, helpers.NewError("Planet not found")
	}

	swapi := client.Swapi{}

	planetAttendanceInFilms, err := swapi.GetPlanetAttendance(planet)

	if err != nil {
		fmt.Printf("swapi not avaiable %s", err)
	}

	planet.Count = planetAttendanceInFilms

	return planet, http.StatusOK, nil
}

func (m *Planet) Create(request *http.Request) (*domain.Planet, int, error) {
	var newPlanet = domain.CreatePlanet{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&newPlanet)

	ok, err := newPlanet.IsValid()

	if !ok || err != nil {
		return nil, http.StatusBadRequest, err
	}

	planet, err := m.Repository.Create(newPlanet)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return planet, http.StatusOK, nil
}

func (m *Planet) Update(request *http.Request) (*domain.Planet, int, error) {
	var planet = domain.Planet{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&planet)
	vars := mux.Vars(request)

	if vars["id"] == "" {
		return nil, http.StatusBadRequest, helpers.NewError("Url Param 'id' is missing")
	}

	isUrlIdValid := govalidator.IsMongoID(vars["id"])

	if !isUrlIdValid {
		return nil, http.StatusBadRequest, helpers.NewError("Not a valid id")
	}

	ok, err := planet.IsValid()

	if !ok || err != nil {
		return nil, http.StatusBadRequest, err
	}

	nPlanet, err := m.Repository.Update(planet)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return nPlanet, http.StatusOK, nil
}

func (m *Planet) Delete(request *http.Request) (int, error) {

	vars := mux.Vars(request)

	if vars["id"] == "" {
		return http.StatusBadRequest, helpers.NewError("Url Param 'id' is missing")
	}

	isUrlIdValid := govalidator.IsMongoID(vars["id"])

	if !isUrlIdValid {
		return http.StatusBadRequest, helpers.NewError("Not a valid id")
	}

	err := m.Repository.Delete(vars["id"])

	if err != nil {
		return http.StatusNotFound, helpers.NewError("Planet not found")
	}

	return http.StatusNoContent, nil
}