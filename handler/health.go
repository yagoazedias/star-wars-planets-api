package handler

import (
	"github.com/unrolled/render"
	"net/http"
)

type Health struct {}

func (*Health) Check(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		_ = formatter.JSON(w, http.StatusOK, struct {
			Status string
		}{"Ok"})
	}
}
