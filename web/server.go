package web

import "github.com/astaxie/beego"


func Server() {
	beego.Run(":19984")
}