package main

import (
	"context"
	pb "daylearn/course/pb/stream/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := pb.NewStreamServiceClient(conn)
	sumCli, err := cli.Sum(context.Background())
	if err != nil {
		panic(err)
	}
	sumCli.Send(&pb.SumRequest{Num: int64(1)})
	sumCli.Send(&pb.SumRequest{Num: int64(2)})
	sumCli.Send(&pb.SumRequest{Num: int64(3)})
	if resp, err := sumCli.CloseAndRecv(); err != nil {
		panic(err)
	} else {
		log.Printf("[D] resp: %d", resp.Result)
	}
}
