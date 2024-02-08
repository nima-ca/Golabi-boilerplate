package bootstrap

import (
	db "Golabi-boilerplate/packages/DB"
	"Golabi-boilerplate/packages/routing"

	"github.com/spf13/viper"
)

func Serve() {

	// Connect to DB
	db.Connect()

	// DOC: initialize the router (Fiber)
	routing.Init()
	routing.RegisterMiddlewares()
	routing.RegisterRoutes()

	// DOC: run the fiber server
	port := viper.GetString("PORT")
	routing.Serve(":" + port)
}
