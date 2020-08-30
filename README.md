# Caddy OpenAPI Validator Example

An example for using the Caddy [OpenAPI Validator](https://github.com/hslatman/caddy-openapi-validator) HTTP handler.

## Description

This repository is an example for the Caddy [OpenAPI Validator](https://github.com/hslatman/caddy-openapi-validator) HTTP handler.
It uses the example `Swagger Petstore` specification.
A minimal (and incomplete) implementation of the API is provided in `internal/petstore/petstore.go`, which only exists for demo purposes.
The `config.json` file is a Caddy configuration file in JSON format.
It configures Caddy to serve the PetStore API with OpenAPI validation, TLS and logging enabled on https://localhost:9443/api.
The example can be run like below:

```bash
# run the main command directly
$ go run cmd/main.go run --config config.json

# compile and run the server
$ go build cmd/main.go
$ ./main run --config config.json
```

The (currently incomplete) API can then be accessed via https://localhost:9443/api/pets/1. 

## TODO

A small and incomplete list of things to add:

* Add more configuration?