package bootstrap

import (
	"Golabi-boilerplate/packages/env"
	"Golabi-boilerplate/packages/routing"

	"github.com/spf13/viper"
)

func Serve() {
	// DOC: load env file
	env.Load()

	// DOC: initialize the router (Fiber)
	routing.Init()
	routing.RegisterMiddlewares()
	routing.RegisterRoutes()

	// DOC: run the fiber server
	port := viper.Get("PORT").(string)
	routing.Serve(":" + port)
}