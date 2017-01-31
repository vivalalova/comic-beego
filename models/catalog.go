package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Catalog struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	title        string
	author       string
	category     string
	description  string
	URL          string
	thumbnailURL string
	_created_at  time.Time
	_updated_at  time.Time
}
