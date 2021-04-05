package main

import (
	"context"
	pb "daylearn/course/pb/stream/proto"
	"io"
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
	facbCli, err := cli.Facb(context.Background(), &pb.FacbRequest{Max: 10})
	if err != nil {
		panic(err)
	}
	for {
		if resp, err := facbCli.Recv(); err != nil {
			if err == io.EOF {
				return
			} else {
				panic(err)
			}
		} else {
			log.Printf("[D] index: %d, facb: %d", resp.Index, resp.Curr)
		}
	}

}
