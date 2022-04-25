# gRPC: SOA booster

> Introduction session to [gRPC](https://grpc.io) as [SOA](https://en.wikipedia.org/wiki/Service-oriented_architecture) booster using [Golang](https://go.dev) and [PHP](https://php.net).

## Table of contents

- [Presentation slides](#presnetation-slides)
- [Golang demo](#golang-demo)
- [PHP demo](#php-demo)

## Presentation slides

The presentation slides related to this repository [can be found here](https://docs.google.com/presentation/d/149Ia9TpyfdponGmLNFyFnPqMib_1-GRX7ruUfuvxnYw).

## Golang demo

- follow [gRPC for Go installation instructions](https://grpc.io/docs/languages/go/quickstart/)

- to run the Go gRPC server, listening on `:50051`

```shell
$ go run ./server
```

- to run the Go gRPC client

```shell
$ go run ./client
```

- if needed, to (re)generate proto Go stubs

```shell
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/textTools.proto
```

## PHP demo

- follow [gRPC for PHP installation instructions](https://grpc.io/docs/languages/php/quickstart/)

- PHP stubs can be found in the [proto/php](proto/php) directory:
    - stubs in the [proto/php/App](proto/php/App) folder
    - example of usage in [proto/php/example.php](proto/php/example.php) (Symfony controller)

- if needed, to (re)generate proto PHP stubs

```shell
protoc --php_out=php \
    --grpc_out=php \
    --plugin=protoc-gen-grpc=~/grpc/cmake/build/grpc_php_plugin \
    proto/textTools.proto
```
