package main

import (
	pb "daylearn/course/pb/stream/proto"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStreamServiceServer
}

//输入流 返回结果
func (*server) Sum(req pb.StreamService_SumServer) (err error) {
	var reqObj *pb.SumRequest

	var sum int64 = 0
	for {
		reqObj, err = req.Recv()
		if err == io.EOF {
			log.Printf("[E] err: %v", err)
			req.SendAndClose(&pb.SumResponse{Result: sum})
			return nil
		} else if err == nil {
			log.Printf("[D] recv: %v", reqObj)
			sum += reqObj.Num
		} else {
			return err
		}
	}
}

//输入正常 返回流
func (*server) Facb(req *pb.FacbRequest, facm pb.StreamService_FacbServer) error {
	if req.Max < 2 {
		facm.Send(&pb.FacbResponse{Index: 1, Curr: 1})
		return nil
	}
	var sum int64 = 1
	var index int32 = 1
	var next int64
	for sum < req.Max {
		facm.Send(&pb.FacbResponse{Index: index, Curr: sum})
		index++
		next, sum = sum, sum+next
	}
	return nil
}

//输入流 返回流
func (*server) Chat(facm pb.StreamService_ChatServer) error {
	for {
		reqObj, err := facm.Recv()
		if err != nil {
			if err == io.EOF || err == nil {
				return nil
			}
			return err
		}
		msg := strings.Replace(reqObj.Msg, "吗", "", -1)
		msg = strings.Replace(msg, "?", "!", -1)
		msg = strings.Replace(msg, "？", "!", -1)
		facm.Send(&pb.ChatResponse{Reply: msg})
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStreamServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
