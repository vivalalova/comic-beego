package models

import (
	"os"

	"fmt"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var (
	session    *mgo.Session
	collection *mgo.Collection
)

func db() {
	var err error

	if mongo := os.Getenv("mongo"); mongo != "" {
		session, err = mgo.Dial(mongo)
		fmt.Println("connect to db:" + mongo)
	} else {
		session, err = mgo.Dial(beego.AppConfig.String("mongo"))
		fmt.Println("connect to db:" + beego.AppConfig.String("mongo"))
	}

	if err != nil {
		if err.Error() == "EOF" {
			session.Refresh()
		}
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
}
