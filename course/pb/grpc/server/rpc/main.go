package main

import (
	"context"
	"log"
	"net"
	"runtime/debug"

	pb "daylearn/course/pb/grpc/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

const (
	port = ":9192"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedHelloServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.String) (*pb.String, error) {
	log.Printf("Received: %v", in.GetValue())
	return &pb.String{Value: "Hello " + in.GetValue()}, nil
}
func (s *server) Hello(ctx context.Context, in *pb.String) (*pb.String, error) {
	log.Printf("Received: %v", in.GetValue())
	// return &pb.String{Value: "Hello this is a tom" + in.GetValue()}, nil
	//标准定义错误状态码
	return nil, status.Error(codes.Canceled, "")
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()
	return handler(ctx, req)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//增加拦截器处理
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			RecoveryInterceptor,
			LoggingInterceptor,
		),
	}
	s := grpc.NewServer(opts...)

	pb.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
