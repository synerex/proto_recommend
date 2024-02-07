# proto_recommend

Created for Joint Research with Hitachi in AY 2023.

## Prerequisites

- [protoc](https://grpc.io/docs/protoc-installation/)
- [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/)

## Build

### for Go lang
```
protoc -I . --go_out=paths=source_relative:. recommend.proto
```
