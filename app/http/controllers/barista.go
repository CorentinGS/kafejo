package controllers

import (
	"github.com/corentings/kafejo/internal/barista"
	"github.com/gofiber/fiber/v2"
)

type BaristaController struct {
	barista.UseCase
}

func (b *BaristaController) GetName(c *fiber.Ctx) error {
	return c.SendString(b.UseCase.GetName())
}
