package main

import (
	"fmt"

	proto "github.com/micro/examples/stream/server/proto"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

func main() {
	service := micro.NewService()
	service.Init()

	// create client
	cl := proto.NewStreamerClient("go.micro.srv.stream", service.Client())

	// create streaming client
	stream, err := cl.Stream(context.Background())
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// send and receive messages for a 10 count
	for j := 0; j < 10; j++ {
		if err := stream.Send(&proto.Request{Count: int64(j)}); err != nil {
			fmt.Println("err:", err)
			return
		}
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println("recv err", err)
			break
		}
		fmt.Printf("Sent msg %v got msg %v\n", j, rsp.Count)
	}

	// close the stream
	if err := stream.Close(); err != nil {
		fmt.Println("stream close err:", err)
	}
}
