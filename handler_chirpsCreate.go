package main

import (
	"encoding/json"
	"net/http"

	"github.com/gio-white/gohttpserver/internal/auth"
	"github.com/gio-white/gohttpserver/internal/database"
)

// func (cfg *apiConfig) handlerChirpsCreate(w http.ResponseWriter, r *http.Request) {
// 	type parameters struct {
// 		Body   string    `json:"body"`
// 		UserID uuid.UUID `json:"user_id"`
// 	}
	
// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
// 		return
// 	}

// 	if len(params.Body) > 140 {
// 		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
// 		return
// 	}

// 	cleanedBody := cleanProfanity(params.Body)

// 	dbChirp, err := cfg.database.CreateChirp(r.Context(), database.CreateChirpParams{
// 		Body:   cleanedBody,
// 		UserID: params.UserID,
// 	})
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp")
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, Chirp{
// 		ID:        dbChirp.ID,
// 		CreatedAt: dbChirp.CreatedAt,
// 		UpdatedAt: dbChirp.UpdatedAt,
// 		Body:      dbChirp.Body,
// 		UserID:    dbChirp.UserID,
// 	})
// }

func (cfg *apiConfig) handlerChirpsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT")
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid JWT")
		return
	}

	params := parameters{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	cleanedBody := cleanProfanity(params.Body)

	dbChirp, err := cfg.database.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   cleanedBody,
		UserID: userID, 
	})

	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserID:    dbChirp.UserID,
	})
}