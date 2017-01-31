package controllers

import (
	"comic-go/models"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CategoryController struct {
	beego.Controller
}

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

func (o *CategoryController) Get() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	Catalog := session.DB("sfacg").C("catalog")
	result := models.Catalog{}
	err = Catalog.Find(bson.M{}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	o.Data["json"] = result
	o.ServeJSON()

	fmt.Println(result)
}
