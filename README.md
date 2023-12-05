# Go-Marketplace backend

## Quickstart

To run app via docker containers, use the command:
```bash
docker compose up
```

If you want to use nightly images, you can pull and run all images from [dockerhub](https://hub.docker.com/):
- [cart-nightly](https://hub.docker.com/repository/docker/almostinf/go-marketplace-cart-nightly/general)
- [user-nightly](https://hub.docker.com/repository/docker/almostinf/go-marketplace-user-nightly/general)
- [product-nightly](https://hub.docker.com/repository/docker/almostinf/go-marketplace-product-nightly/general)
- [order-nightly](https://hub.docker.com/repository/docker/almostinf/go-marketplace-order-nightly/general)
- [gateway-nightly](https://hub.docker.com/repository/docker/almostinf/go-marketplace-gateway-nightly/general)

To run unit tests:
```bash
make test
```

## Architecture

The backend features an intricate microservices architecture, with each service having its own database and seamless interaction through APIs. For detailed APIs, check the `/proto` folder

The server component uses the [grpc protocol](https://github.com/grpc/grpc-go), enabling exclusive communication among microservices through grpc. Users can choose between grpc and the Restful API via the [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) mechanism

- The user service is vital for storing and modifying user information

- The cart service manages cart details and items, addressing prolonged product storage with a worker. The worker, accessing Redis, cleans up the cart and returns products periodically

- The product service is key for managing product information, ensuring product deletions reflect in associated cart items. Also the service stores information about the discount in Redis with a user-defined life time

- The order service oversees order data, allowing status changes and user order cancellations within 24 hours. Upon order or part deletion, all products are returned

- The gateway service acts as a user facade and authorizes requests, directing them to the necessary microservices for streamlined system functionality

## Docs

All project documentation is in the `/docs` folder, e.g. swagger documentation, ER diagrams and so on

Also, when you run the application under the path `/api/v1/swagger`, you can see the swagger ui and play around with the application api

## Dependencies

[gprc protocol](https://github.com/grpc/grpc-go) was used for server and client side, also from grpc ecosystem I used [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) and [protoc](https://grpc.io/docs/protoc-installation/)

**For unit tests:**
- [testify](https://github.com/stretchr/testify)
- [gomock](https://github.com/golang/mock)

**To work with the databases:**
- PostgreSQL:
  - [pgx](https://github.com/jackc/pgx)
  - [goose](https://github.com/pressly/goose)
  - [squirrel](https://github.com/Masterminds/squirrel)
- Redis:
  - [go-redis](https://github.com/redis/go-redis)

**For working with configuration files:** [cleanenv](https://github.com/ilyakaznacheev/cleanenv)

**Logger:** [zerolog](https://github.com/rs/zerolog)

**Linters**:
- [Smart Imports](https://github.com/pav5000/smartimports)
- [Golang-ci lint](https://golangci-lint.run/)

**Swagger**: [swaggerui](https://github.com/flowchartsman/swaggerui)

**RBAC**: [gorbac](https://github.com/mikespook/gorbac)

**Validation**: [validator](https://github.com/go-playground/validator)

**JWT**: [jwt-go](https://github.com/golang-jwt/jwt)

## License
- MIT license ([LICENSE-MIT](https://github.com/seanmonstar/httparse/blob/master/LICENSE-MIT) or [https://opensource.org/licenses/MIT](https://opensource.org/licenses/MIT))