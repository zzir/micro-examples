# gRPC Config Server

## Get Started

### config server

```bash
cd config-srv
go run main.go
```

### client

```bash
go run client/main.go
```

### change files

cd to config-srv/conf, change **micro.yml** file any values.

eg.

```bash
micro:
  name: Micro
  version: 1.0.0
  hi: hello
```

to 

```bash
micro:
  name: Micro
  version: 1.0.0
  hi: I am fine
```

then the client print:

```bash
2019/04/28 10:57:15 Watch changes: {"hi":"I am fine","name":"Micro","version":"1.0.0"}
``` 