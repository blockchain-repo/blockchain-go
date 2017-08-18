package controllers

import (
	"unichain-go/common"
	"unichain-go/config"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["json"] = common.Serialize(config.Config)
	this.ServeJSON()
}