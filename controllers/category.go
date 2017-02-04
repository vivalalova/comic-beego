package controllers

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

type CategoryController struct {
	BaseController
}

func (o *CategoryController) Get() {
	var results []string
	err := Catalog.Find(bson.M{}).Distinct("category", &results)

	if err != nil {
		log.Fatal(err)
	}

	o.Data["json"] = results
	o.ServeJSON()
}
