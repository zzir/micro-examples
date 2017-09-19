# API

This repo contains examples for serving microservices via the micro api.

The [micro api](https://github.com/micro/micro/tree/master/api) is an API gateway which serves HTTP and routes to RPC based services. 
In the micro ecosystem we logically separate concerns via architecture and tooling. Read more on buiding an API layer of services 
in the [architecture blog post](https://micro.mu/blog/2016/04/18/micro-architecture.html).

The micro api by default serves the namespace go.micro.api. Our service names include this plus a unique name e.g go.micro.api.example. 
You can change the namespace via the flag `--namespace=`.

The micro api has a number of different handlers which lets you define what kind of API services you want. See examples below. The handler 
can be set via the flag `--handler=`.

## Contents

- default - an api using the default micro api request handler
- proxy - an api using the http proxy handler
- rpc - an api using standard go-micro rpc
- meta - an api which specifies its api endpoints

## Deps

Service discovery is required for all services. Default is Consul or MDNS. You can also use plugins from 
[micro/plugins](https://github.com/micro/go-plugins).

### MDNS

Use the flag `--registry=mdns`

### Consul

```
brew install consul
consul agent -dev
```

