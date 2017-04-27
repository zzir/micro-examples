# Stream

This is an example of a streaming service and client

## Contents

- server - is the service
- client - is the client

## Prereqs

Micro services need a discovery system so they can find each other. Micro uses consul by default but 
its easily swapped out with etcd, kubernetes, or various other systems. We'll run consul for convenience.

Install consul
```shell
brew install consul
```

Alternative instructions - [https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

Run Consul

```shell
consul agent -dev
```

## Run the example

Run the service

```shell
go run server/main.go
```

Run the client

```shell
go run client/main.go
```

And that's all there is to it.
