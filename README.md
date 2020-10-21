# Caddy OpenAPI Validator Example

An example for using the Caddy [OpenAPI Validator](https://github.com/hslatman/caddy-openapi-validator).

## Description

This repository is an example for using the Caddy [OpenAPI Validator](https://github.com/hslatman/caddy-openapi-validator) HTTP handler.
It's based on the example (expanded) `Swagger Petstore` specification.
A minimal (and incomplete) implementation of the API is provided in `internal/petstore/petstore.go`, which only exists for demo purposes.
We've also included an example of the expanded PetStore API based on [deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen/tree/master/examples/petstore-expanded).
The `config.json` file is a Caddy configuration file in JSON format.
It configures Caddy to serve the PetStore API with OpenAPI validation, TLS and logging enabled on https://localhost:9443/api.
The Expanded PetStore example configuration is in `expanded.json` and also includes all of the default configuration values for the [OpenAPI Validator](https://github.com/hslatman/caddy-openapi-validator).

The example can be started as shown below:

```bash
# run the main command directly
$ go run cmd/main.go run --config=config.json

# compile and run the server
$ go build cmd/main.go
$ ./main run --config=config.json

# run with the expanded pet store example
$ go run cmd/main.go run --config=expanded.json
```

The API can then be accessed in your browser of choice: https://localhost:9443/api/pets/1. 
