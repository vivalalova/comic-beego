package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var (
	session    *mgo.Session
	collection *mgo.Collection
)

func db() {
	var err error
	session, err = mgo.Dial(beego.AppConfig.String("mongo"))

	if err != nil {
		if err.Error() == "EOF" {
			session.Refresh()
		}
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
}
