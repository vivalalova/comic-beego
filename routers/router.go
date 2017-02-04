// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"comic-go/controllers"

	"github.com/astaxie/beego"
)

func init() {

	// beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	// beego.AutoRouter(&controllers.CategoryController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/catalog", &controllers.CatalogController{})
	beego.Router("/chapter", &controllers.ChapterController{})
	beego.Router("/page", &controllers.PageController{})
}
