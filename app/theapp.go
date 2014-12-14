package app

import (
	_ "github.com/astaxie/beego/config"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ghstahl/pingbeego/configuration"
	"github.com/ghstahl/pingbeego/filters"
	"github.com/astaxie/beego/context"
	"code.google.com/p/log4go"
	log "github.com/inconshreveable/log15"
)

type LogContext struct {

}
func (self *LogContext) String() string {
	return "well hello"
}
var TheLogContext = &LogContext{}

var TheApp = &App{}
type App struct {

}

func (self *App) Initialize() {
	self.initializeViewStarts()
	self.initializeKeyValues()
	self.initializeAuth()
	self.initializeFilters()
}

func (self *App) initializeAuth() {
}

var FilterTest = func(ctx *context.Context) {
	fmt.Printf(">>>>>>FilterTest Enter\n")
}


func (self *App) initializeFilters() {
	log4go.Trace("initializeFilters: %s (%d)", "herb", 3)
	filters.TheFilterConfigs.Load()
	beego.InsertFilter("/", beego.BeforeRouter, FilterTest)//matches only the / route
	beego.InsertFilter("*", beego.BeforeRouter, FilterTest)// matches everything else, but the / route
}

func (self *App) initializeKeyValues() {
	configuration.TheKeyValueConfigs.Load();
}
func Tag() string{
	return "well hello"
}
func (self *App) initializeViewStarts() {
	configuration.TheViewStartConfigs.Load();
}

func (self *App) Run() {
	log4go.LoadConfiguration("log4go.xml")
	srvlog := log.New("module", "app/server")
	// all log messages can have key/value context
	srvlog.Warn("abnormal conn rate", "rate", TheLogContext, "low", 2, "high", 3)

	fmt.Println("Running App")
	self.Initialize()
	beego.Run()
}

