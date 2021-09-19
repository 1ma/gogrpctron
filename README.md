## gogrpctron

An exploratory Go GRPC service built from scratch.

## Prerequisites

* GNU Make
* [Go toolchain 1.17+](https://golang.org/dl/) with the `$GOPATH/bin` directory in `$PATH`
* [`protoc`](https://github.com/protocolbuffers/protobuf/releases/) in `$PATH`
* `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
* `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
* Google's "well known protobuf types" in `/usr/local/include/google` (included in the `protoc` zip file)
