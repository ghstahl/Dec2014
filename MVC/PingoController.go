package MVC

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/ghstahl/pingbeego/configuration"
	"github.com/ghstahl/pingbeego/filters"
	log "github.com/inconshreveable/log15"

	"github.com/astaxie/beego/context"
)



type PingoController struct {
	beego.Controller
 	theLogger log.Logger
}

func (this *PingoController) Logger() log.Logger{
	return this.theLogger
}

func (this *PingoController) PrepareLoggingContext(ctx *context.Context) {
	requestId := filters.GetCurrentRequestId(ctx);
	this.theLogger = log.New(filters.WellKnown.RequestId, requestId)
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


