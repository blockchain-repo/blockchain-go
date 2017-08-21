package web

import (
	_ "unichain-go/web/routers"

	"github.com/astaxie/beego"
)


func Server() {
	beego.BConfig.CopyRequestBody =true
	beego.Run("localhost:19984")
}