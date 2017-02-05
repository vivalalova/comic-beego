package controllers

import (
	"comic-go/models"
	"strconv"

	"github.com/astaxie/beego"
)

type params struct {
	limit int
	skip  int
	sort  string
}

// just for inherit
type BaseController struct {
	beego.Controller
}

var (
	Catalog = models.Catalog{}.Shared()
	Chapter = models.Chapter{}.Shared()
)

func (controller *BaseController) parms() params {
	sort := controller.GetString("sort")

	if sort == "" {
		sort = "-_updated_at"
	}

	return params{
		limit: controller.limit(),
		skip:  controller.skip(),
		sort:  sort,
	}
}

func (controller *BaseController) limit() int {
	limitString := controller.GetString("limit")
	limit, _ := strconv.Atoi(limitString)

	if limit == 0 || limit > 100 {
		limit = 30
	}

	return limit
}

func (controller *BaseController) skip() int {
	skipString := controller.GetString("skip")
	skip, _ := strconv.Atoi(skipString)
	return skip
}
