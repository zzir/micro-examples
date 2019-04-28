package main

import (
	"github.com/micro/go-config"
	grpcConfig "github.com/micro/go-config/source/grpc"
	"github.com/micro/go-log"
)

type Micro struct {
	Demo
}

type Demo struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Hi      string `json:"hi,omitempty"`
	Age     int    `json:"age,omitempty"`
}

func main() {

	source := grpcConfig.NewSource(
		grpcConfig.WithAddress("127.0.0.1:8600"),
		grpcConfig.WithPath("/micro"),
	)

	conf := config.NewConfig()
	err := conf.Load(source)
	if err != nil {
		log.Fatal(err)
	}

	configs := &Micro{}
	err = conf.Scan(configs)
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("Read configs: %s", string(conf.Bytes()))

	watcher, err := conf.Watch()
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("Watch changes ...")

	for {
		v, err := watcher.Next()
		if err != nil {
			log.Fatal(err)
		}

		log.Logf("Watch changes: %v", string(v.Bytes()))
	}
}
