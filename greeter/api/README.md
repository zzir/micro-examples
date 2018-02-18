# API

This directory showcases API services which sit behind the micro api and serve a public facing API

## Services

- [api.go](api.go) - runs an RPC api with [api.Request](https://github.com/micro/go-api/blob/master/proto/api.proto#L11L18) and [api.Response](https://github.com/micro/go-api/blob/master/proto/api.proto#L21L25)
- [gin.go](gin.go) - using gin server. set micro api with --handler=proxy
- [rest.go](rest.go) - using go-restful. set micro api with --handler=proxy
- [rpc.go](rpc.go) - using RPC. set micro api with --handler=rpc
- 

