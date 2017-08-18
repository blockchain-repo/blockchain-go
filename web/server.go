package web

import (
	_ "unichain-go/web/routers"

	"github.com/astaxie/beego"
)


func Server() {
	beego.Run("localhost:19984")
}