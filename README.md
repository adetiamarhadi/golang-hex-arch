generate schema

```sh
protoc --go_out=internal\adapters\framework\left\grpc --proto_path=internal\adapters\framework\left\grpc\proto\ internal\adapters\framework\left\grpc\proto\number_msg.proto
```

generate service

```sh
protoc --go-grpc_out=internal\adapters\framework\left\grpc --proto_path=internal\adapters\framework\left\grpc\proto internal\adapters\framework\left\grpc\proto\arithmetic_svc.proto
```
