package controllers

import (
	"comic-go/models"



	"gopkg.in/mgo.v2/bson"
)

type CatalogController struct {
	BaseController
}

func (controller *CatalogController) queries() bson.M {

	queries := bson.M{}

	if category := controller.GetString("category"); len(category) > 0 {
		queries["category"] = category
	}

	if title := controller.GetString("title"); len(title) > 0 {
		queries["title"] = title
	}

	return queries
}

func (controller *CatalogController) Get() {

	id := controller.Ctx.Input.Param(":id")

	if id == "" {
		controller.find()
	} else {
		controller.findOne(id)
	}
}

func (controller *CatalogController) find() {
	var results []models.Catalog

	query := controller.queries()

	err := Catalog.Find(query).
		Skip(controller.parms().skip).
		Limit(controller.parms().limit).
		Sort(controller.parms().sort).
		All(&results)

	if err != nil {
		controller.Abort(err.Error())
	}

	controller.Data["json"] = results
	controller.ServeJSON()
}

func (controller *CatalogController) findOne(id string) {
	//find One
	result := []models.Catalog{}

	err := Catalog.Find(bson.M{"ID": id}).Limit(1).All(&result)

	if err != nil {
		controller.Abort(err.Error())
	}



	controller.Data["json"] = result
	controller.ServeJSON()
}
