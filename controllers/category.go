package controllers

import (
	"comic-go/models"
	"log"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type CategoryController struct {
	beego.Controller
}

var (
	Catalog = models.Catalog{}.Shared()
	results []string
)

func (o *CategoryController) Get() {

	err := Catalog.Find(bson.M{}).Distinct("category", &results)
	// defer Catalog.Database.Session.Close()

	if err != nil {
		log.Fatal(err)
		err = nil
	}

	o.Data["json"] = results
	o.ServeJSON()

}
