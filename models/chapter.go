package models

import (
	"time"

	"gopkg.in/mgo.v2"
)

type Chapter struct {
	ID        string    `bson:"_id,omitempty" json:"_id"`
	CatalogID string    `bson:"catalogID" json:"catalogID"`
	Title     string    `bson:"title" json:"title"`
	URL       string    `bson:"URL" json:"url"`
	Pages     []string  `bson:"pages" json:"pages"`
	CreatedAt time.Time `bson:"_created_at" json:"_created_at"`
	UpdatedAt time.Time `bson:"_updated_at" json:"_updated_at"`
}

// type ChapterNext struct {
// 	Chapter
// 	Prev Chapter
// 	Next Chapter
// }

func (chapter Chapter) Shared() *mgo.Collection {
	db()
	return session.DB("sfacg").C("chapter")
}
