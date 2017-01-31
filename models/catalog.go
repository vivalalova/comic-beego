package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Catalog struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Title        string        `bson: "title"`
	Author       string
	Category     string
	Description  string
	URL          string
	ThumbnailURL string
	_created_at  time.Time
	_updated_at  time.Time
}
