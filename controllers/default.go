package controllers

import (
	"github.com/ghstahl/Dec2014/MVC"
	"github.com/ghstahl/pingbeego/filters"
	"time"
)

type MainController struct {
	MVC.PingoController
}

func (this *MainController) Prepare() {
	this.PrepareLayout("Primary")
	this.PrepareLoggingContext(this.Ctx)
}

func (this *MainController) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter

	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		this.Logger().Warn("abnormal conn rate", "rate", 3, "low", 2, "high", 3)

	}



	this.Data[filters.WellKnown.RequestId] = filters.GetCurrentRequestId(this.Ctx)
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "ghstahl@gmail.com"
	this.Data["Request"] = r
	this.Data["ResponseWriter"] = w
	this.TplNames = "index.tpl"
}
