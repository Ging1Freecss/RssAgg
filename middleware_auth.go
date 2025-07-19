package main

import (
	"net/http"

	"github.com/Ging1Freecss/RssAgg/internal/auth"
	db "github.com/Ging1Freecss/RssAgg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, db.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPiKey(r.Header)

		if err != nil {
			responseWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			responseWithError(w, http.StatusNotFound, "Couldn't get user")
			return
		}

		handler(w,r,user)
	}
}
