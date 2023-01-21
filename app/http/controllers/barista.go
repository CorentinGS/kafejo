package controllers

import (
	"github.com/corentings/kafejo/internal/barista"
)

type BaristaController struct {
	barista.UseCase
}

func (b *BaristaController) GetName(c *fiber.Ctx) error {

}
