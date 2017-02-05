package controllers

import "gopkg.in/mgo.v2/bson"

type CategoryController struct {
	BaseController
}

func (controller *CategoryController) Get() {
	var results []string
	err := Catalog.Find(bson.M{}).Distinct("category", &results)

	if err != nil {
		controller.Abort(err.Error())
	}

	controller.Data["json"] = results
	controller.ServeJSON()
}
