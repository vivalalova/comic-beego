package controllers

import (
	"comic-go/models"

	"gopkg.in/mgo.v2/bson"
)

type PageController struct {
	BaseController
}

func (controller *PageController) Get() {

	if id := controller.Ctx.Input.Param(":id"); id != "" {
		result := models.Chapter{}
		err := Chapter.Find(bson.M{"_id": id}).One(&result)

		if err != nil {
			controller.Abort(err.Error())
		} else if result.ID == "" {
			controller.CustomAbort(404, "not found")
		}
		controller.Data["json"] = result.Pages
		controller.ServeJSON()
	} else {
		controller.CustomAbort(404, "chapter id is nessary")
	}
}
