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

### With Go CLI

#### Run http server
```shell
 go run cmd/main.go http 
```

#### Run grpc server
```shell
 go run cmd/main.go grpc 
```
