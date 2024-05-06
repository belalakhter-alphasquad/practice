package main

import (
	"fmt"

	api "github.com/belalakhter-alphasquad/practice/internal/api"
	DB "github.com/belalakhter-alphasquad/practice/internal/db"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	fmt.Println("Server is Starting")
    dbclient := DB.NewClient()
	handler := api.NewHandler(dbclient)
	api.SetupRoutes(app,handler)

	err := app.Listen("localhost:8080")
	if err != nil {
		fmt.Printf("Error started server %v", err)
	}
}
