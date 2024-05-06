package api

import (
	"context"
	"fmt"
	"time"

	"github.com/belalakhter-alphasquad/practice/internal/db"
	fiber "github.com/gofiber/fiber/v2"
)

type Handler struct {
	dbclient *db.Connection
}


func NewHandler( dbclient *db.Connection) *Handler{
	return &Handler{dbclient:  dbclient}
}

func (h *Handler) PingHandler(c *fiber.Ctx) error{
 return c.Status(fiber.StatusOK).JSON(map[string]string{"Result:" : "Pong"})
}


func (h *Handler) GetUserHandler(c *fiber.Ctx) error {
    ctx,cancel := context.WithTimeout(c.Context(),15 * time.Second)
	defer cancel()
	users,err := db.GetUserRepository(ctx,h.dbclient)
      if(err != nil ){
		fmt.Printf("Error fetching User %v",err)
	  }
    return c.Status(fiber.StatusOK).JSON(users)
}