package main

import (
	"context"
	"errors"

	proto "github.com/jenmud/protoc-gen-go-nats-grpc-adaptor/example"
	"google.golang.org/protobuf/types/known/structpb"
)

type DemoService struct {
	proto.UnimplementedGreeterServer
}

func (s *DemoService) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hi " + req.GetName()}, nil
}

func (s *DemoService) SayHelloAgain(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return nil, errors.New("some random example error")
}

func (s *DemoService) SayGoodbye(ctx context.Context, req *proto.SayGoodbyeRequest) (*proto.SayGoodbyeReply, error) {
	return &proto.SayGoodbyeReply{Message: "Later " + req.GetName()}, nil
}

func (s *DemoService) SaveMetadata(ctx context.Context, req *structpb.Struct) (*structpb.Struct, error) {
	m := req.AsMap()
	m["saved"] = true

	updated, err := structpb.NewStruct(m)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
