package models

import "gopkg.in/mgo.v2"

var (
	session    *mgo.Session
	collection *mgo.Collection
)

func db() {
	var err error
	session, err = mgo.Dial("localhost")

	if err != nil {
		if err.Error() == "EOF" {
			session.Refresh()
		}
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
}
