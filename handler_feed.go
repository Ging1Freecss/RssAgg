package main

import (
	"encoding/json"
	"net/http"
	"time"

	db "github.com/Ging1Freecss/RssAgg/internal/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (apiCfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID: pgtype.UUID{
			Bytes: uuid.New(),
			Valid: true,
		},
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		UserID: user.ID,
		Name:   params.Name,
		Url:    params.URL,
	})

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	respondWithJSON(w, http.StatusOK, dbFeedToFeed(feed))
}
