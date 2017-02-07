package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Catalog collection
type Catalog struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Identifier   string        `bson:"ID" json:"ID"`
	Title        string        `bson:"title" json:"title"`
	Author       string        `bson:"author" json:"author"`
	Category     string        `bson:"category" json:"category"`
	Description  string        `bson:"description" json:"description"`
	URL          string        `bson:"URL" json:"url"`
	ThumbnailURL string        `bson:"thumbnailURL" json:"thumbnailURL"`
	CreatedAt    time.Time     `bson:"_created_at" json:"_created_at"`
	UpdatedAt    time.Time     `bson:"_updated_at" json:"_updated_at"`
}

func (catalog Catalog) Shared() *mgo.Collection {
	db()
	return session.DB("sfacg").C("catalog")
}
