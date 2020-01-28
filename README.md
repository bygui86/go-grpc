
# go-grpc
gRPC example in Golang

## Services

- [greeting service (gRPC client)](greeting-service)
- [hello service (gRPC server)](hello-service)

---

## Run

1. Compile protobufs
	```shell
	protoc -I go/proto/ go/proto/ --go_out=plugins=grpc:go/domain
	```

2. Start server (hello-service)
	```shell
	cd hello-service
	GO111MODULE=on go run main.go
	```

3. In another shell, start client (greeting-service)
	```shell
	cd greeting-service
	GO111MODULE=on go run main.go
	```

---

## Polyglot test

This repo can be used for a polyglot test together with [java-grpc](https://github.com/bygui86/java-grpc)

### Go client --> Java server

1. Start server (hello-service)
	```shell
	cd java-grpc/hello-service
	mvnw clean spring-boot:run
	```

2. In another shell, start client (greeting-service)
	```shell
	cd go-grpc/greeting-service
	GO111MODULE=on go run main.go
	```

### Java client --> Go server

1. Start server (hello-service)
	```shell
	cd go-grpc/hello-service
	GO111MODULE=on go run main.go
	```

2. In another shell, start client (greeting-service)
	```shell
	cd java-grpc/greeting-service
	mvnw clean spring-boot:run
	```

---

## TODO list

- [ ] testing

---

## Links

- https://grpc.io/docs/quickstart/go.html
