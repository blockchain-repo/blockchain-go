package controllers

import (
	"encoding/json"

	"unichain-go/common"
	"unichain-go/config"
	"unichain-go/core"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type TXController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Body([]byte(common.Serialize(config.Config)))
}

func (this *TXController) Post() {
	//get json
	var requestParamMap map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &requestParamMap)
	//add key 'assign' and 'assign_timestamp'
	//insert to backlog
	core.WriteTransactionToBacklog(requestParamMap)
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Body([]byte(common.Serialize(requestParamMap)))
}
