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

//just for this api
type chapter struct {
	models.Chapter
	prev models.Chapter
	next models.Chapter
}

func (controller *ChapterController) findOne(catalogID string, id string) chapter {

	var chapters []models.Chapter

	err := Chapter.
		Find(bson.M{"catalogID": catalogID}).
		Select(bson.M{"pages": 0}).
		Sort("-title").
		All(&chapters)

	if err != nil {
		controller.Abort(err.Error())
	}

	var result chapter

	fmt.Println(id)

	err = Chapter.Find(bson.M{"_id": id}).One(&result)

	if err != nil {
		controller.Abort(err.Error())
	}

	fmt.Println(result)

	if result.ID == "" {
		controller.CustomAbort(404, "not found")
	}

	for index, element := range chapters {
		if element.ID == result.ID {
			if index > 0 {
				result.prev = chapters[index-1]
			}

			if index < len(chapters) {
				result.next = chapters[index+1]
			}
		}
	}

	return result
}
