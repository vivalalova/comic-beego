package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Catalog struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	Title        string        `bson:"title" json:"title"`
	Author       string        `bson:"author" json:"author"`
	Category     string        `bson:"category" json:"category"`
	Description  string        `bson:"description" json:"description"`
	URL          string        `bson:"URL" json:"url"`
	ThumbnailURL string        `bson:"thumbnailURL" json:"thumbnailURL"`
	CreatedAt    time.Time     `bson:"_created_at" json:"createdAt"`
	UpdatedAt    time.Time     `bson:"_updated_at" json:"updatedAt"`
}

var (
	session    *mgo.Session
	collection *mgo.Collection
)

func (catalog Catalog) Shared() *mgo.Collection {
	if session == nil {
		var err error
		session, err = mgo.Dial("localhost")

		if err != nil {
			panic(err)
		}

		session.SetMode(mgo.Monotonic, true)
		collection = session.DB("sfacg").C("catalog")
	}

	return collection
}
