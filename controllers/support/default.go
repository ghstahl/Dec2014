package support

import (
	"html/template"
	"github.com/ghstahl/Dec2014/MVC"
)



type SupportController struct {
	MVC.PingoController
}

func (this *SupportController) Prepare() {
	this.PrepareLayout("Primary")
}


func (this *SupportController) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	this.Data["xsrfdata"]=template.HTML(this.XsrfFormHtml())
	this.Data["xsrf_token"] = this.XsrfToken()

	this.Data["Website"] = "beego.me Support"
	this.Data["Email"] = "doggy@gmail.com"
	this.Data["Request"] = r
	this.Data["ResponseWriter"] = w
	this.TplNames = "support/index.tpl"
}


