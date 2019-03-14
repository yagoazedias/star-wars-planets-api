package config

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"github.com/yagoazedias/star-wars-planets-api/handler"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRouters(mx, formatter)

	n.UseHandler(mx)

	return n
}

func initRouters(mx *mux.Router, formatter *render.Render) {

	health := handler.Health{}
	planet := handler.Planet{}

	mx.HandleFunc("/health", health.Check(formatter)).Methods("GET")
	mx.HandleFunc("/planet", planet.Search(formatter)).Methods("GET")
	mx.HandleFunc("/planet", planet.Create(formatter)).Methods("POST")
	mx.HandleFunc("/planet/id/{id}", planet.Lookup(formatter)).Methods("GET")
	mx.HandleFunc("/planet/id/{id}", planet.Update(formatter)).Methods("PUT")
}