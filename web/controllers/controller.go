package controllers

import (
	"encoding/json"

	"unichain-go/common"
	"unichain-go/config"
	"unichain-go/core"
	"unichain-go/log"
	"unichain-go/models"

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
	var tx models.Transaction
	json.Unmarshal(this.Ctx.Input.RequestBody, &tx)
	log.Debug("Api receive tx", tx.Id)
	flag := core.ValidateTransaction(tx)
	if flag {
		core.WriteTransactionToBacklog(tx)
		this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		this.Ctx.Output.Status = 202
		this.Ctx.Output.Body([]byte(common.Serialize(tx)))
	} else {
		this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		this.Ctx.Output.Status = 400
		this.Ctx.Output.Body([]byte(common.Serialize(tx)))
	}
}
