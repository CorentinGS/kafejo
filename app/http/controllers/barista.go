package controllers

import (
	"github.com/corentings/kafejo/internal/barista"
	"github.com/gofiber/fiber/v2"
)

type BaristaController struct {
	barista.IUseCase
}

func NewBaristaController(useCase barista.IUseCase) BaristaController {
	return BaristaController{IUseCase: useCase}
}

func (b *BaristaController) GetName(c *fiber.Ctx) error {
	return c.SendString(b.IUseCase.GetName())
}
