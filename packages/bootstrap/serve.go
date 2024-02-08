package bootstrap

import (
	"Golabi-boilerplate/packages/routing"

	"github.com/spf13/viper"
)

func Serve() {
	// DOC: initialize the router (Fiber)
	routing.Init()
	routing.RegisterMiddlewares()
	routing.RegisterRoutes()

	// DOC: run the fiber server
	port := viper.GetString("PORT")
	routing.Serve(":" + port)
}
