package controllers

import (
	"comic-go/models"
	"log"

	"gopkg.in/mgo.v2/bson"
)

type CatalogController struct {
	BaseController
}

func (this *CatalogController) queries() bson.M {
	category := this.GetString("category")
	title := this.GetString("title")

	queries := bson.M{}

	if len(category) > 0 {
		queries["category"] = category
	}

	if len(title) > 0 {
		queries["title"] = title
	}

	return queries
}

func (this *CatalogController) Get() {

	id := this.Ctx.Input.Param(":id")

	if id == "" {
		var results []models.Catalog

		query := this.queries()

		err := Catalog.Find(query).
			Skip(this.parms().skip).
			Limit(this.parms().limit).
			Sort(this.parms().sort).
			All(&results)

		if err != nil {
			log.Fatal(err)
		}

		this.Data["json"] = results
		this.ServeJSON()
	} else {
		result := models.Catalog{}

		err := Catalog.Find(bson.M{"ID": id}).One(&result)

		if err != nil {
			this.CustomAbort(400, err.Error())
		}

		this.Data["json"] = result
		this.ServeJSON()
	}
}
