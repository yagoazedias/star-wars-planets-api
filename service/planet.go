package service

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/yagoazedias/star-wars-planets-api/domain"
	"github.com/yagoazedias/star-wars-planets-api/helpers"
	"github.com/yagoazedias/star-wars-planets-api/repository"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Planet struct {
	Repository repository.Planet
}

func (m *Planet) Search() ([]domain.Planet, int, error)  {
	planets, err := m.Repository.Search()

	if err != nil {
		return nil, http.StatusInternalServerError, helpers.NewError("An unexpected error has occurred")
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
		return nil, http.StatusInternalServerError, helpers.NewError("An unexpected error has occurred")
	}

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