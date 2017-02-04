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
	CreatedAt    time.Time     `bson:"_created_at" json:"createdAt"`
	UpdatedAt    time.Time     `bson:"_updated_at" json:"updatedAt"`
}

var (
	session    *mgo.Session
	collection *mgo.Collection
)

func (catalog Catalog) Shared() *mgo.Collection {

	var err error
	session, err = mgo.Dial("localhost")

	if err != nil {
		if err.Error() == "EOF" {
			session.Refresh()
		}
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	collection = session.DB("sfacg").C("catalog")

	return collection
}
