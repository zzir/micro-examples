package main

import (
	"log"

	"github.com/micro/go-micro"
	"golang.org/x/net/context"

	proto "github.com/micro/examples/stream/server/proto"
)

type Streamer struct{}

func (e *Streamer) Stream(ctx context.Context, stream proto.Streamer_StreamStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Got msg %v", req.Count)
		if err := stream.Send(&proto.Response{Count: req.Count}); err != nil {
			return err
		}
	}
}

func main() {
	// new service
	service := micro.NewService(
		micro.Name("go.micro.srv.stream"),
	)

	// Init command line
	service.Init()

	// Register Handler
	proto.RegisterStreamerHandler(service.Server(), new(Streamer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
