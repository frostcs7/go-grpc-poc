package main

import (
	pb "github.com/frostcs7/go-grpc-poc/gen/protos"
	"golang.org/x/net/context"
	"reflect"
	"testing"
)

func TestServer_SayHello(t *testing.T) {
	type fields struct {
		UnsafeHelloWorldServer pb.UnsafeHelloWorldServer
	}
	type args struct {
		context context.Context
		in      *pb.HelloRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "world",
			args: args{in: &pb.HelloRequest{Name: "world"}},
			want: "Hello world",
		},
		{
			name: "123",
			args: args{in: &pb.HelloRequest{Name: "123"}},
			want: "Hello 123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				UnsafeHelloWorldServer: tt.fields.UnsafeHelloWorldServer,
			}
			got, err := s.SayHello(tt.args.context, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Message, tt.want) {
				t.Errorf("SayHello() got = %v, want %v", got, tt.want)
			}
		})
	}
}
