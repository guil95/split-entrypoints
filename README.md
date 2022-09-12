# split-entrypoints

## How to use

### With bin file
```shell
 go build -o app cmd/main.go 
```

#### Run http server
```shell
 ./app http
```

#### Run grpc server
```shell
 ./app grpc
```

#### Run consumers
```shell
 ./app consumers
```

### With Go CLI

#### Run http server
```shell
 go run cmd/main.go http 
```

#### Run grpc server
```shell
 go run cmd/main.go grpc 
```

#### Run consumers
```shell
 go run cmd/main.go consumers 
```

### With docker

#### Run all the entrypoints
```shell
docker-compose up --build -d
```
