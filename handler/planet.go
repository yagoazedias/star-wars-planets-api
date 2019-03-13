package handler

import (
	"github.com/unrolled/render"
	"net/http"
)

type Planet struct {}

func (*Planet) Search(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		_ = formatter.JSON(w, http.StatusOK, []struct {
			Name string
			Weather string
			Soil string
		}{{"Tatooine", "hot", "sandy"}, {"Tatooine", "hot", "sandy"}, {"Tatooine", "hot", "sandy"}})
	}
}

func (*Planet) Create(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		_ = formatter.JSON(w, http.StatusCreated, struct {
			Name string
			Weather string
			Soil string
		}{"Tatooine", "hot", "sandy"})
	}
}