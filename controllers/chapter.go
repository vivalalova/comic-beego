package controllers

import (
	"comic-go/models"

	"gopkg.in/mgo.v2/bson"
)

type ChapterController struct {
	BaseController
}

func (controller *ChapterController) Get() {

	catalogID := controller.Ctx.Input.Param(":catalogID")

	if id := controller.Ctx.Input.Param(":id"); id != "" {
		controller.Data["json"] = controller.findOne(catalogID, id)
	} else {
		controller.Data["json"] = controller.find(catalogID)
	}

	controller.ServeJSON()
}

func (controller *ChapterController) find(catalogID string) []models.Chapter {
	var chapters []models.Chapter
	err := Chapter.
		Find(bson.M{"catalogID": catalogID}).
		Select(bson.M{"pages": 0}).
		Sort("-title").
		All(&chapters)

	if err != nil {
		controller.Abort(err.Error())
	}

	return chapters
}

func (controller *ChapterController) findOne(catalogID string, id string) models.Chapter {

	var chapters []models.Chapter
	err := Chapter.
		Find(bson.M{"catalogID": catalogID}).
		Select(bson.M{"pages": 0}).
		Sort("-title").
		All(&chapters)

	if err != nil {
		controller.Abort(err.Error())
	}

	return models.Chapter{}
}
