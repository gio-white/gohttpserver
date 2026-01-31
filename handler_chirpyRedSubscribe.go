package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gio-white/gohttpserver/internal/auth"
	"github.com/google/uuid"
)


type WebhookRequest struct {
	Event string `json:"event"`
	Data  struct {
		UserID string `json:"user_id"`
	} `json:"data"`
}

func (cfg *apiConfig) handlerPolkaWebhook(w http.ResponseWriter, r *http.Request) {
	key, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't retrieve the API key")
	}
	if key != cfg.polkaKey {
		respondWithJSON(w, http.StatusUnauthorized, "Invalid API key")
	}

	decoder := json.NewDecoder(r.Body)
	req := WebhookRequest{}
	err = decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}
	
	if req.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	userID, err := uuid.Parse(req.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	err = cfg.database.SetChirpyRedByID(r.Context(), userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "User not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Couldn't update user")
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}