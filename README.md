# A Microservice Architecture Project

## Prerequisites

You need **Make**, **Docker** and **docker-compose** installed because everything is running in Docker container and we are using a `Makefile` to simplify the whole process.

## Getting Started

```sh
make up
```

It will build Docker image and run monolith and microservices versions.

### Services addresses

Monolith: `http://localhost:8090/`

Orders microservice: `http://localhost:8070/`

Shop microservice: `http://localhost:8071/`

Payments microservice: no public API (you can export ports in `docker-compose.yml` if you need)

For available methods, please check interfaces layer in source code and tests `tests/acceptance_test.go`.

## Running tests

First of all you must run services

```sh
make up
```

Then you can run all tests by using in another terminal:

```sh
make docker-test
```

If you want to test only monolith version:

```sh
make docker-test-monolith
```

or microservices:

```sh
make docker-test-microservices
```
