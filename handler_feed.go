package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Toha22BSK/rssagg/gen/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}
