package handler

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/unrolled/render"
	"github.com/yagoazedias/star-wars-planets-api/service"
	"net/http"
)

type Planet struct {
	Service service.Planet
}

func (h *Planet) Search(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {

		offset := request.URL.Query().Get("offset")
		limit := request.URL.Query().Get("limit")

		planets, status, err := h.Service.Search(offset, limit)


		if err != nil {
			_ = formatter.JSON(w, status, bson.M{
				"message": fmt.Sprintf("%q", err.Error()),
			})
		} else {
			_ = formatter.JSON(w, status, planets)
		}
	}
}

func (h *Planet) Lookup(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		planet, status, err := h.Service.Lookup(request)

		if err != nil {
			_ = formatter.JSON(w, status, bson.M{
				"message": fmt.Sprintf("%q", err.Error()),
			})
		} else {
			_ = formatter.JSON(w, status, planet)
		}
	}
}

func (h *Planet) Create(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {

		newPlanet, status, err := h.Service.Create(request)

		if err != nil {
			_ = formatter.JSON(w, status, bson.M{
				"message": fmt.Sprintf("%q", err.Error()),
			})
		} else {
			_ = formatter.JSON(w, status, newPlanet)
		}
	}
}

func (h *Planet) Update(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {

		newPlanet, status, err := h.Service.Update(request)

		if err != nil {
			_ = formatter.JSON(w, status, bson.M{
				"message": fmt.Sprintf("%q", err.Error()),
			})
		} else {
			_ = formatter.JSON(w, status, newPlanet)
		}
	}
}