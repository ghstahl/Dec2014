package routers

import (
	"github.com/astaxie/beego"
	"github.com/ghstahl/Dec2014/controllers"
	"github.com/ghstahl/Dec2014/controllers/support"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/support", &support.SupportController{})
}
