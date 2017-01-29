package controllers

import "github.com/astaxie/beego"

// Operations about category
type CategoryController struct {
	beego.Controller
}

func (o *CategoryController) GetAll() {
	// obs := models.GetAll()
	// o.Data["json"] = obs
	// o.ServeJSON()
	o.Ctx.WriteString("hello world")
}

func (o *CategoryController) Get() {
	o.Ctx.WriteString("hello world2")
}
