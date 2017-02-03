package controllers

import (
	"comic-go/models"
	"log"

	"gopkg.in/mgo.v2/bson"
)

type CatalogController struct {
	BaseController
}

// type CatalogParams struct {
// 	limit    int
// 	skip     int
// 	category string
// 	title    string
// }

// func (this *string) toCatalogParams() CatalogParams {
// 	return CatalogParams{}
// }

func (this *CatalogController) Get() {

	// category := this.GetString("category")
	// title := this.GetString("title")
	// fmt.Println(category)
	// fmt.Println(title)

	var results []models.Catalog

	query := bson.M{}
	// query["category"] = "辛巴达的冒险"

	err := Catalog.Find(query).
		Skip(this.skip()).
		Limit(this.limit()).
		All(&results)

	if err != nil {
		log.Fatal(err)
	}

	this.Data["json"] = results
	this.ServeJSON()

}
