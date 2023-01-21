package barista

type Handlers interface {
	// GetName returns the name of the barista.
	GetName() string
}

type UseCase interface {
	// GetName returns the name of the barista.
	GetName() string
}

type PgRepository interface {
	// GetName returns the name of the barista.
	GetName() string
}

type RedisRepository interface {
	// GetName returns the name of the barista.
	GetName() string
}
