package subscriber

import (
	"github.com/micro/go-log"

	"context"
	example "github.com/micro/examples/template/fnc/proto/example"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}
