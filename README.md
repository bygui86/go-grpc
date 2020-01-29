
# go-grpc
gRPC example in Golang

## Services

- [hello service (gRPC server)](hello-service)
- [greeting service (gRPC client)](greeting-service)

---

## Codebase

### Prerequisites

1. Compile protobufs
	```shell
	protoc --proto_path=./proto/ --go_out=plugins=grpc:domain ./proto/*
	```

### Build

1. Server (hello-service)
	```shell
	go build -o grpc-server ./hello-service
	```

2. Client (greeting-service)
	```shell
	go build -o grpc-client ./greeting-service
	```

### Run

1. Start server (hello-service)
	```shell
	GO111MODULE=on go run ./hello-service/main.go
	```

2. In another shell, start client (greeting-service)
	```shell
	GO111MODULE=on go run ./greeting-service/main.go
	```

---

## Docker

### Build

1. Server (hello-service)
	```shell
	docker build -t grpc/hello-service -f hello.Dockerfile .
	```

2. Client (greeting-service)
	```shell
	docker build -t grpc/greeting-service -f greeting.Dockerfile .
	```

### Run

1. Server (hello-service)
	```shell
	docker run -ti --rm --name hello-service -p 50051:50051 grpc/hello-service
	```

2. Client (greeting-service)
	```shell
	docker run -ti --rm --name greeting-service grpc/greeting-service
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
	GO111MODULE=on go run ./greeting-service/main.go
	```

### Java client --> Go server

1. Start server (hello-service)
	```shell
	GO111MODULE=on go run ./hello-service/main.go
	```

2. In another shell, start client (greeting-service)
	```shell
	cd java-grpc/greeting-service
	mvnw clean spring-boot:run
	```

---

## TODO list

- [ ] dockerfiles
- [ ] kubernetes probes
- [ ] kubernetes manifests
- [ ] testing

---

## Links

- https://grpc.io/docs/quickstart/go.html
