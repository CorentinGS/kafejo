package internal

import (
	"github.com/corentings/kafejo/app/http/controllers"
	"github.com/corentings/kafejo/infrastructures"
	"github.com/corentings/kafejo/internal/barista"
	"sync"
)

type ServiceContainer interface {
	GetBarista() controllers.BaristaController
}

type kernel struct{}

func (k kernel) GetBarista() controllers.BaristaController {
	barisataRepo := barista.BaristaPgRepository{DBConn: infrastructures.GetDBConn()}
	baristaUsecase := barista.BaristaUsecase{PgRepository: &barisataRepo}
	baristaController := controllers.BaristaController{UseCase: &baristaUsecase}

	return baristaController
}

var (
	k             *kernel   // kernel is the service container
	containerOnce sync.Once // containerOnce is used to ensure that the service container is only initialized once
)

// GetServiceContainer returns the service container
func GetServiceContainer() ServiceContainer {
	containerOnce.Do(func() {
		k = &kernel{}
	})
	return k
}
