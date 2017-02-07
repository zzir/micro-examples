package main

import (
	"log"

	"github.com/micro/examples/template/srv/handler"
	"github.com/micro/examples/template/srv/subscriber"
	"github.com/micro/go-micro"

	example "github.com/micro/examples/template/srv/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.template"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	example.RegisterSubscriber("topic.go.micro.srv.template", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	example.RegisterSubscriber("topic.go.micro.srv.template", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
