package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService()
	service.Init()
	client := service.Client()

	request := client.NewJsonRequest("greeter", "Greeter.Hello", "john")
	var response string

	if err := client.Call(context.TODO(), request, &response); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
