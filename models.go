package main

import (
	db "github.com/Ging1Freecss/RssAgg/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	Name      string           `json:"name"`
	ApiKey    string           `json:"api_key"`
}

func dbUserToUser(dbUser db.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	Name      string           `json:"name"`
	Url       string           `json:"url"`
	UserID    pgtype.UUID      `json:"user_id"`
}

func dbFeedToFeed(feed db.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}
