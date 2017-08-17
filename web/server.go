package web

import "github.com/astaxie/beego"


func Server() {
	beego.Run("localhost:19984")
}