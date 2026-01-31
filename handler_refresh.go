package main

import (
	"net/http"
	"time"

	"github.com/gio-white/gohttpserver/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	refreshTokenStr, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find token")
		return
	}

	dbRefreshToken, err := cfg.database.GetRefreshToken(r.Context(), refreshTokenStr)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid refresh token")
		return
	}

	if dbRefreshToken.ExpiresAt.Before(time.Now().UTC()) {
		respondWithError(w, http.StatusUnauthorized, "Refresh token expired")
		return
	}
	if dbRefreshToken.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "Refresh token revoked")
		return
	}

	accessToken, err := auth.MakeJWT(dbRefreshToken.UserID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create access token")
		return
	}

	respondWithJSON(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{
		Token: accessToken,
	})
}