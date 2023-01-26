create:
#	protoc -I ./protos \
#       --go_out ./gen/protos --go_opt paths=source_relative \
#       --go-grpc_out ./gen/protos --go-grpc_opt paths=source_relative \
#       ./protos/helloworld.proto
	buf generate

clean:
	rm gen/protos/*.go