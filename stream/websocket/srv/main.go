package main

import (
	"log"

	proto "github.com/micro/examples/stream/websocket/srv/proto"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Streamer struct {
	stream proto.Streamer_ServerStreamStream
}

func (s *Streamer) ServerStream(ctx context.Context, req *proto.Request, stream proto.Streamer_ServerStreamStream) error {
	log.Printf("Got msg %v", req.Count)

	// Stream i responses to server
	for i := 1; i <= int(req.Count); i++ {
		if err := stream.Send(&proto.Response{Count: int64(i)}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.stream"),
	)

	service.Init()

	proto.RegisterStreamerHandler(service.Server(), new(Streamer))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
