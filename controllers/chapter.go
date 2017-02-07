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

func (controller *ChapterController) findOne(catalogID string, id string) models.ChapterWithNext {

	var chapters []models.Chapter

	err := Chapter.
		Find(bson.M{"catalogID": catalogID}).
		Select(bson.M{"pages": 0}).
		Sort("-title").
		All(&chapters)

	if err != nil {
		controller.Abort(err.Error())
	}

	var chapter models.Chapter
	var result models.ChapterWithNext

	err = Chapter.Find(bson.M{"_id": id}).One(&chapter)

	if err != nil {
		controller.Abort(err.Error())
	} else if chapter.ID == "" {
		controller.CustomAbort(404, "not found")
	} else {
		result.Chapter = chapter

		for index, element := range chapters {
			if element.ID == result.ID {
				if index > 0 {
					result.Prev = &chapters[index-1]
				}

				if index < len(chapters) {
					result.Next = &chapters[index+1]
				}
			}
		}
	}

	return result
}
