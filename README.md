# go-grpc-poc
This is poc for golang grpc

## To generate the proto service definitions
`make create`

This command will trigger the buf generate to create the files in the ./gen/protos folder

## To run
`go run server/server.go`

The server is under server/server.go.

## How to test
To list all the service definitions:

`grpCurl --plaintext localhost:50051 list`

To send request:

`grpCurl --plaintext -d '{"name": "sandeep"}' localhost:50051 protos.HelloWorld/SayHello`

Using buf curl:

`buf curl --protocol grpc --http2-prior-knowledge http://localhost:50051/protos.HelloWorld/SayHello`

### Unit test
`go test ./...`

## Buf useful commands

`buf lint`

`buf breaking --against '.git#branch=main'`

`buf format`