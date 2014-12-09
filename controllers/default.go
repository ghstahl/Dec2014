package controllers

import (
	"github.com/ghstahl/Dec2014/MVC"
	"github.com/ghstahl/pingbeego/filters"
)

type MainController struct {
	MVC.PingoController
}

func (this *MainController) Prepare() {
	this.PrepareLayout("Primary")
}

func (this *MainController) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter


	this.Data[filters.WellKnown.RequestId] = filters.GetCurrentRequestId(this.Ctx)
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "ghstahl@gmail.com"
	this.Data["Request"] = r
	this.Data["ResponseWriter"] = w
	this.TplNames = "index.tpl"
}
