package http

import (
	"github.com/corentings/kafejo/internal"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(router *fiber.Router) {
	baristaController := internal.GetServiceContainer().GetBarista()
	(*router).Add("GET", "/barista", baristaController.GetName)
}
