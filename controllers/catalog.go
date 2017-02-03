package controllers

import (
	"comic-go/models"
	"log"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type CatalogController struct {
	beego.Controller
}

func (this *CatalogController) Get() {

	this.GetStrings()

	var results []models.Catalog

	err := Catalog.Find(bson.M{}).Limit(30).All(&results)

	if err != nil {
		log.Fatal(err)
	}

	this.Data["json"] = results
	this.ServeJSON()
}
