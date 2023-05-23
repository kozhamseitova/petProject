package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/pets", app.createPetHandler)
	router.HandlerFunc(http.MethodGet, "/v1/pets", app.showPetsHandler)

	router.HandlerFunc(http.MethodGet, "/v1/pets/:id", app.showPetHandler)

	return router
}
