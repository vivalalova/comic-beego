package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (self *BaseController) limit() int {
	limitString := self.GetString("limit")
	limit, _ := strconv.Atoi(limitString)

	if limit == 0 {
		limit = 30
	}
	return limit
}

func (self *BaseController) skip() int {
	skipString := self.GetString("skip")
	skip, _ := strconv.Atoi(skipString)
	return skip
}
