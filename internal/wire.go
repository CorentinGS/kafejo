//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/corentings/kafejo/app/http/controllers"
	"github.com/corentings/kafejo/infrastructures"
	"github.com/corentings/kafejo/internal/barista"
	"github.com/google/wire"
)

func InitializeBarista() controllers.BaristaController {
	wire.Build(controllers.NewBaristaController, barista.NewUseCase, barista.NewPgRepository, infrastructures.GetDBConn)
	return controllers.BaristaController{}
}
