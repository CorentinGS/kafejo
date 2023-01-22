package controllers

import (
	"github.com/corentings/kafejo/internal/kliento"
	"github.com/gofiber/fiber/v2"
)

type KlientoController struct {
	kliento.IUseCase
}

func NewKlientoController(useCase kliento.IUseCase) KlientoController {
	return KlientoController{IUseCase: useCase}
}

func (k *KlientoController) GetName(c *fiber.Ctx) error {
	return c.SendString(k.IUseCase.GetName())
}

func (k *KlientoController) SetName(c *fiber.Ctx) error {
	name := c.Params("name")
	err := k.IUseCase.SetName(name)
	if err != nil {
		return err
	}
	return c.SendString("Name set to " + name)
}
