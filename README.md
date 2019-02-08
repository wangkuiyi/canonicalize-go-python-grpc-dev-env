This repo demonstrates how to write a gRPC server in Go and a client in Python.

## Canonical Development Environment

A secondary motivation of this project is to show how to create a canonical development environment for Go and Python programmers using Docker.  

### Building in Container

We will build a Docker image that contains development tools:

1. The Python interpreter
1. The Go compiler
1. The protobuf compiler
1. The protobuf to Go compiler extension
1. The protobuf to Python compiler extension

### Editing on Host

When we use this Docker image for daily development work, the source code relies in the host computer instead of the container.  The source ocde includes this repo and all its dependencies, for example, the Go package `google.golang.org/grpc`.  This allows us to run our favorite editors (Emacs, VIM, Eclipse, etc) on the host.  Please free to rely on editors addons to analyze the source code for auto-completion.


## How to Build

Because this repo contains Go code, please make sure that you have the directory structure required by Go.  On my laptop computer, I have

```bash
export GOPATH=$HOME/go
```

You could have your `$GOPATH` pointing to any directory you like.

Given `$GOPATH$` set, we could git clone the source code of our project and all its dependencies, including `google.golang.org/grpc`, by running:

```bash
go get github.com/wangkuiyi/multi-stream-grpc
```

To build this demo, we need the protobuf compiler, Go compiler, Python interpreter, gRPC extension to protobuf compiler.  To ease the installation and configuration of these tools, I provide a Dockerfile to install them into a Docker image. To build the Docker image

```bash
cd $GOPATH/src/github.com/wangkuiyi/multi-stream-grpc
docker build -t grpc .
```

To run the container, we need to map the `$GOPATH` directory on the host into the `/go` directory in the container, because the Dockerfile configures `/go` as the `$GOPATH` in the container:

```bash
docker run --rm -it \
    -v $PWD:/go/src/github.com/wangkuiyi/multi-stream-grpc \
    -w /go/src/github.com/wangkuiyi/multi-stream-grpc \
    grpc bash
```

Now, in the container, we can compile the `sqlflow.proto` in this repo into the Go source code:

```bash
protoc -I proto proto/sqlflow.proto --go_out=plugin=grpc:proto
```

Similarly, we can compile it into Python:

```bash
python -m grpc_tools.protoc -I proto --python_out=client --grpc_python_out=client sqlflow.proto
```

Please be aware that the Go toolchain requires that the generated Go source files be in the same directory as the `.proto` file, which is a separate directory than the server source code, whereas the Python convention is to put generated files with the client source code.

To build the Go server:

```bash
cd server
go get -u ./...
go generate
go install
```

where the `go get -u ./...` retrieves and updates Go dependencies of our server, `go generate` invokes the `protoc` command to translate `proto/sqlflow.proto` into `proto/sqlflow.pb.go`, and `go install` builds the server into `$GOPATH/bin/server`.

To run the Go server:

```bash
$GOPATH/bin/server &
```

To run the Python client:

```bash
cd ../client
python client_test.py
```
