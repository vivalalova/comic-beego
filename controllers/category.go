package controllers

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CategoryController struct {
	beego.Controller
}

type Person struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string
	Phone string
}

func (o *CategoryController) Get() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	People := session.DB("test").C("people")

	result := Person{}
	err = People.Find(bson.M{}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result)

	o.Data["json"] = result
	o.ServeJSON()

	// result = Catalog{}
	// println(result)
}
