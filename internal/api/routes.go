package api

import (
	fiber "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App,handler *Handler)  {

api := app.Group("/api")
api.Get("/ping",handler.PingHandler)
api.Get("/GetUser", handler.GetUserHandler)

}
