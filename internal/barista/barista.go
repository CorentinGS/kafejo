package barista

type Handlers interface {
	// GetName returns the name of the barista.
	GetName() string
}

type IUseCase interface {
	// GetName returns the name of the barista.
	GetName() string
}

type IPgRepository interface {
	// GetName returns the name of the barista.
	GetName() string
}

type RedisRepository interface {
	// GetName returns the name of the barista.
	GetName() string
}
