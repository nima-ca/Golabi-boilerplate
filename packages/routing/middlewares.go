package routing

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

// DOC: Register all global middlewares
func RegisterMiddlewares() {
	router := GetRouter()

	// Handle Panics
	router.Use(recover.New())

	// Add security headers
	router.Use(helmet.New())

	// Add CORS policies
	router.Use(cors.New(cors.Config{
		AllowOrigins:     viper.GetString("ORIGINS"),
		AllowCredentials: true,
	}))

	// Logger Middleware
	router.Use(logger.New(logger.Config{
		Format:     "${pid}  |  [${time}]  |  ${status}  |  ${latency}  |  ${method} - ${path}\n",
		TimeFormat: "15:04:05 02-Jan-2006",
		TimeZone:   "Asia/Tehran", // Change timezone based on your location
	}))

}
