package main

import (
	"encoding/json"
	"net/http"
)


func (apiCfg *apiConfig) handlerPostUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	dbUser, err := apiCfg.database.CreateUser(r.Context(), params.Email)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
        return
    }

	respondWithJSON(w, http.StatusCreated, User{
        ID:        dbUser.ID,	
        CreatedAt: dbUser.CreatedAt,
        UpdatedAt: dbUser.UpdatedAt,
        Email:     dbUser.Email,
    })
}