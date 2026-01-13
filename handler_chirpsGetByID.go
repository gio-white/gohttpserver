package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerChirpsGetByID(w http.ResponseWriter, r *http.Request){
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID")
		return
	}

	dbChirp, err := apiCfg.database.GetChirpsByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found")
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserID:    dbChirp.UserID,
	})
}