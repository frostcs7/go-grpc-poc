create:
	protoc --go_out="./gen/protos" protos/helloworld.proto

clean:
	rm gen/proto/*.go