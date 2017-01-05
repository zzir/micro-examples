# Micro Examples

This is a repository for micro examples. Feel free to contribute.

## Deps

All services require service discovery. The default is Consul or MDNS.

### Consul 

```
brew install consul
consul agent -dev
```

### MDNS

Use flag `--registry=mdns`

## Here

- [greeter](greeter) - A complete greeter example (includes python, ruby examples)
- [sidecar](sidecar) - Greeter service using the sidecar with multiple languages
- [booking](booking) - A booking.com demo application
- [plugins](plugins) - How to use plugins
- [template](template) - Api, web and srv service templates generated with `micro new`

## External

- [auth-srv](https://github.com/micro/auth-srv) - An Oauth2 authentication service
- [geo-srv](https://github.com/micro/geo-srv) - Geolocation tracking service using hailocab/go-geoindex
- [geo-web](https://github.com/micro/geo-web) - Web demo for the geo srv
- [geo-api](https://github.com/micro/geo-api) - API for the geo srv
- [discovery-srv](https://github.com/micro/discovery-srv) - A discovery in the micro platform
- [geocode-srv](https://github.com/micro/geocode-srv) - A geocoding service using the Google Geocoding API
- [hailo-srv](https://github.com/micro/hailo-srv) - A service for the hailo taxi service developer api
- [monitor-srv](https://github.com/micro/monitor-srv) - A monitoring service for Micro services
- [place-srv](https://github.com/micro/place-srv) - A microservice to store and retrieve places (includes Google Place Search API)
- [slack-srv](https://github.com/micro/slack-srv) - The slack bot API as a go-micro RPC service
- [trace-srv](https://github.com/micro/trace-srv) - A distributed tracing microservice in the realm of dapper, zipkin, etc
- [twitter-srv](https://github.com/micro/twitter-srv) - A microservice for the twitter API
- [user-srv](https://github.com/micro/user-srv)	- A microservice for user management and authentication

