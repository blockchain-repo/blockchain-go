package routers

import (
	"github.com/astaxie/beego"
	"unichain-go/web/controllers"
)


func init() {
	beego.Router("/", &controllers.MainController{})
}