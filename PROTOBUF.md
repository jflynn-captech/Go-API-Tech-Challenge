# Install protobuf tool


install protobuf via brew

```bash
brew install protobuf
```

Go dependencies
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

# Generating types and services from schema.

```bash
protoc --go_out=data/ --go-grpc_out=data/ data.proto
```
