package main

import (
	_ "github.com/ghstahl/Dec2014/routers"
	"github.com/ghstahl/Dec2014/app"
)

func main() {
	app.TheApp.Run()
}

