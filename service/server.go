package service

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"net/http"
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
	mx.HandleFunc("/health", healthCheckHandler(formatter)).Methods("GET")
}

func healthCheckHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		_ = formatter.JSON(w, http.StatusOK, struct {
			Status string
		}{"Ok"})
	}
}