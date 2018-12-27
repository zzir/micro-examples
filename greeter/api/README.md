# API

This directory showcases API services which sit behind the micro api and serve a public facing API

## Services

- [**api**](api.go) - runs an RPC api with [api.Request](https://github.com/micro/go-api/blob/master/proto/api.proto#L11L18) and [api.Response](https://github.com/micro/go-api/blob/master/proto/api.proto#L21L25)
- [**gin**](gin) - using gin server
- [**rest**](rest) - using go-restful
- [**rpc**](rpc) - using RPC
