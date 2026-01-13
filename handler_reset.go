package main

import "net/http"

func (apiCfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if apiCfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}

	err := apiCfg.database.ResetUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't reset database")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database reset successfully"))
	
	apiCfg.fileserverHits.Swap(0)
}