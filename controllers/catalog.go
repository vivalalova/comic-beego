package controllers

import (
	"comic-go/models"
	"log"
	"strconv"

	"fmt"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type CatalogController struct {
	beego.Controller
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
	limitString := this.GetString("limit")
	limit, _ := strconv.Atoi(limitString)

	if limit == 0 {
		limit = 30
	}

	skipString := this.GetString("skip")
	skip, _ := strconv.Atoi(skipString)

	category := this.GetString("category")
	title := this.GetString("title")
	fmt.Println(category)
	fmt.Println(title)

	var results []models.Catalog

	query := bson.M{}
	// query["category"] = "辛巴达的冒险"

	err := Catalog.Find(query).Skip(skip).Limit(limit).All(&results)

	if err != nil {
		log.Fatal(err)
	}

	this.Data["json"] = results
	this.ServeJSON()

}
