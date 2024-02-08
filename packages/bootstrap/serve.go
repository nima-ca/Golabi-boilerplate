package bootstrap

import (
	db "Golabi-boilerplate/packages/db"
	"Golabi-boilerplate/packages/routing"
	"Golabi-boilerplate/packages/validators"

	"github.com/spf13/viper"
)

func Serve() {
	// Connect to DB
	db.Connect()

	// Register custom validators
	validators.RegisterValidators()

	// DOC: initialize the router (Fiber)
	routing.Init()
	routing.RegisterMiddlewares()
	routing.RegisterRoutes()

	// DOC: run the fiber server
	port := viper.GetString("PORT")
	routing.Serve(":" + port)
}
