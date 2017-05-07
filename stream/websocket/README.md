# Websocket Stream

This is an example of a streaming service and a web client that forwards the stream to a websocket.

## Contents

- srv - is the service
- web - is the web client

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

The web client runs behind the `micro web` reverse proxy

Run micro web

``` shell
micro web
```

## Run the example

Run the service

```shell
go run srv/main.go
```

Run the web client

```shell
cd web # must be in the web directory to serve static files.
go run main.go
```

Visit http://localhost:8082/stream and send a request!
