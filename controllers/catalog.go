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

	queries := bson.M{}

	if category := this.GetString("category"); len(category) > 0 {
		queries["category"] = category
	}

	if title := this.GetString("title"); len(title) > 0 {
		queries["title"] = title
	}

	return queries
}

func (this *CatalogController) Get() {

	id := this.Ctx.Input.Param(":id")

	if id == "" {
		this.find()
	} else {
		this.findOne(id)
	}
}

func (this *CatalogController) find() {
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
}

func (this *CatalogController) findOne(id string) {
	//find One
	result := models.Catalog{}

	err := Catalog.Find(bson.M{"ID": id}).One(&result)

	if err != nil {
		this.CustomAbort(400, err.Error())
	}

	this.Data["json"] = result
	this.ServeJSON()
}
