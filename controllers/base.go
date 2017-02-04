package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)

type params struct {
	limit int
	skip  int
	sort  string
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) parms() params {
	sort := this.GetString("sort")

	if sort == "hot" {
		sort = "hot"
	} else {
		sort = "-_updated_at"
	}

	return params{
		limit: this.limit(),
		skip:  this.skip(),
		sort:  sort,
	}
}

func (this *BaseController) limit() int {
	limitString := this.GetString("limit")
	limit, _ := strconv.Atoi(limitString)

	if limit == 0 || limit > 100 {
		limit = 30
	}

	return limit
}

func (this *BaseController) skip() int {
	skipString := this.GetString("skip")
	skip, _ := strconv.Atoi(skipString)
	return skip
}
