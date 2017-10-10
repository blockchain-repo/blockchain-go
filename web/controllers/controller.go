package controllers

import (
	"encoding/json"

	"unichain-go/common"
	"unichain-go/config"
	"unichain-go/core"
	"unichain-go/log"

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
	var txMap map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &txMap)
	//TODO validate tx
	log.Debug("Api receive tx",txMap["id"])
	core.WriteTransactionToBacklog(txMap)
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Body([]byte(common.Serialize(txMap)))
}
