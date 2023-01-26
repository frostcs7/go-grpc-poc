# go-grpc-poc
This is poc for golang grpc

## To generate the proto service definitions
`make create`
This command will trigger the protoc to create the files in the ./gen/protos folder

## To run
`go run server/server.go`
The server is under server/server.go.

## How to test
To list all the service definitions:
`grpCurl --plaintext localhost:50051 list`

To send request:
`grpCurl --plaintext -d '{"name": "sandeep"}' localhost:50051 protos.HelloWorld/SayHello`

