# Kafejo coffee shop template

Kafejo is a simple and clean backend template in Go.

I created this template to help me start new projects faster. 
It's a simple template with a few features that I use in most of my projects. 

I had trouble finding a good architecture for my projects, so I decided to create this template to help me start new projects faster.


## Features

### Database

- [x] PostgreSQL support with [GORM](https://gorm.io/)
- [ ] MongoDB support
- [ ] Cassandra support with [gocql](https://github.com/gocql/gocql)
- [ ] BBolt support with [bbolt](https://github.com/etcd-io/bbolt)

### Authentication

- [ ] JWT authentication
- [ ] OAuth2 authentication
- [ ] Bcrypt password hashing
- [ ] Argon2 password hashing
- [ ] Sessions management

### Caching

- [x] Redis support with [go-redis](https://github.com/go-redis/redis) and [dragonfly](https://github.com/dragonflydb/dragonfly/blob/main/docs/quick-start/README.md)
- [ ] Custom memory cache

### Logging

- [x] Zerolog support with [zerolog](https://github.com/rs/zerolog)

### Testing

- [ ] Testing with [Testify](https://github.com/stretchr/testify)
- [ ] Mocking with [Mockery](https://github.com/vektra/mockery)
- [ ] Memory leak detection with [uber-go/goleak](https://github.com/uber-go/goleak)
- [ ] Lint with [golangci-lint](https://github.com/golangci/golangci-lint)

### Documentation

- [ ] Swagger support with [swaggo](https://github.com/swaggo/swag)
- [ ] GoDoc support with [godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)

### Deployment

- [ ] Docker support
- [ ] Podman support

### Other

- [ ] Configuration with [Viper](https://github.com/spf13/viper)
- [ ] RabbitMQ support with [amqp](https://github.com/rabbitmq/amqp091-go)
- [ ] Web framework with [Fiber](https://github.com/gofiber/fiber)
- [ ] Prometheus 
- [ ] Jaeger
- [ ] Grafana


## Directory structure

```
.
├── app
│   └── http
│       ├── app.go
│       ├── controllers
│       │   ├── barista.go
│       │   └── kliento.go
│       ├── middleware.go
│       └── routes.go
├── config
│   ├── config.go
│   └── const.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── domain
│   ├── barista.go
│   ├── coffee.go
│   └── kliento.go
├── go.mod
├── go.sum
├── infrastructures
│   ├── postgresql.go
│   ├── rabbitmq.go
│   └── redis.go
├── internal
│   ├── barista
│   │   ├── barista.go
│   │   ├── pgRepository.go
│   │   └── usecase.go
│   ├── containers.go
│   ├── kliento
│   │   ├── kliento.go
│   │   ├── redisRepository.go
│   │   └── usecase.go
│   ├── wire_gen.go
│   └── wire.go
├── LICENSE
├── main.go
├── pkg
│   ├── json
│   │   ├── go-json.go
│   │   ├── json.go
│   │   └── sonic.go
│   ├── logger
│   │   └── logger.go
│   ├── utils
│   └── views
└── README.md
```

## License

This project is licensed under the ISC License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

- [go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [go-realworld-clean](https://github.com/err0r500/go-realworld-clean)
- [wire](https://github.com/google/wire)
- [disgo](https://github.com/switchupcb/disgo)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request
6. Wait for review
7. Enjoy :coffee:

## Disclaimer

This project should be used for educational purposes only. 
The author is not responsible for any damage caused by this project.

I tried to make this project as simple as possible, but it still has a lot of room for improvement.
It's not perfect, but it's a good starting point, and it's a good example of how to structure your project.

## Contact

If you have any questions, feel free to [contact me](https://corentings.vercel.app/links/).




