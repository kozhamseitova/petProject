package main

import (
	"fmt"
	"net/http"
	"petProject.kozhamseitova.net/internal/data"
	"time"
)

func (app *application) createPetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new pet")
}

func (app *application) showPetsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "show all pets")
}

func (app *application) showPetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	pet := data.Pet{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "Bob",
		BirthDate: time.Date(2021, 8, 15, 0, 0, 0, 0, time.Local),
		SpeciesID: 1,
	}

	err = app.writeJSON(w, http.StatusOK, pet, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
