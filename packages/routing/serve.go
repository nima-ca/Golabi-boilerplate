package routing

import (
	"log"
)

// DOC: it bootstraps the Fiber on given port
func Serve(port string) {
	router := GetRouter()

	log.Fatal(router.Listen(port))
}
