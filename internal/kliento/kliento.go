package kliento

type Handlers interface {
	// GetName returns the name of the kliento.
	GetName() string
}

type UseCase interface {
	// GetName returns the name of the kliento.
	GetName() string
}

type PgRepository interface {
	// GetName returns the name of the kliento.
	GetName() string
}

type RedisRepository interface {
	// GetName returns the name of the kliento.
	GetName() string
}
