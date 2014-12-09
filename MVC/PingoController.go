package MVC

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/ghstahl/pingbeego/configuration"
)



type PingoController struct {
	beego.Controller
}

func (this *PingoController) PrepareLayout(viewStartKey string) {
	fmt.Println("Start func (c *PingoController) Prepare()")

	var config = configuration.TheViewStartConfigs.ViewStartConfigMap[viewStartKey]

	this.LayoutSections = make(map[string]string)

	this.LayoutSections["SharedHead"] = config.SharedHead
	this.LayoutSections["Header"] = config.Header
	this.LayoutSections["Footer"] = config.Footer
	this.LayoutSections["HtmlHead"] = config.HtmlHead
	this.Layout = config.Layout

	fmt.Println("END func (c *PingoController)  Prepare()")
}


