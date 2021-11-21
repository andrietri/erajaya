package main

import (
	"erajaya/app"
	"erajaya/config"
)

func init() {
	config.LoadENV()
}

func main() {
	app.NewApp().Run(config.ENV.AppPort)
}
