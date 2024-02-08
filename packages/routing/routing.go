package routing

import (
	userRoutes "Golabi-boilerplate/app/user/routes"

	"github.com/gofiber/fiber/v2"
)

// DOC: it creates a router and assign it to GlobalRouter variable
func Init() {
	GlobalRouter = fiber.New()
}

// DOC: it returns the Global Router
func GetRouter() *fiber.App {
	return GlobalRouter
}

// DOC: it registers all the routes in different modules
func RegisterRoutes() {
	router := GetRouter()

	// Serve public folder
	// For testing you can go to this path in your browser
	// http://localhost:3000/public/sample.jpg
	router.Static("/public", "./public")

	// Register your routes here:
	userRoutes.Register(router)
}
