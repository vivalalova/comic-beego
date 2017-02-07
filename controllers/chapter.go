package controllers

import (
	"comic-go/models"

	"fmt"

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

	fmt.Println(catalogID)

	return chapters
}

func (controller *ChapterController) findOne(catalogID string, id string) models.ChapterNext {

	var chapters []models.Chapter

	err := Chapter.
		Find(bson.M{"catalogID": catalogID}).
		Select(bson.M{"pages": 0}).
		Sort("-title").
		All(&chapters)

	if err != nil {
		controller.Abort(err.Error())
	}

	var result models.ChapterNext

	err = Chapter.Find(bson.M{"_id": id}).One(&result)

	fmt.Println(result)

	if err != nil {
		controller.Abort(err.Error())
	} else if result.ID == "" {
		controller.CustomAbort(404, "not found")
	}

	for index, element := range chapters {
		if element.ID == result.ID {
			if index > 0 {
				result.Prev = chapters[index-1]
			}

			if index < len(chapters) {
				result.Next = chapters[index+1]
			}
		}
	}

	return result
}
