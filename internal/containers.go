package internal

import "sync"

type ServiceContainer interface {
	// GetKliento returns the kliento service
	GetKliento()
}

type kernel struct{}

func (k kernel) GetKliento() {
	//TODO implement me
	panic("implement me")
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
