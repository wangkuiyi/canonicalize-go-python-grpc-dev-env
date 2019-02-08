This repo demonstrates how to write a gRPC server in Go and a client in Python.

## How to Build

Because this repo contains Go code, please make sure that you have the directory structure required by Go.  On my laptop computer, I would do

```bash
mkdir -p ~/work/src/github.com/wangkuiyi # or any other directory
```

To retrieve the source code into the correct directory:

```bash
cd ~/work/src/github.com/wangkuiyi
git clone https://github.com/wangkuiyi/multi-stream-grpc
```

To build this demo, we need the protobuf compiler, Go compiler, Python interpreter, gRPC extension to protobuf compiler.  To ease the installation and configuration of these tools, I provide a Dockerfile to install them into a Docker image. To build the Docker image

```bash
cd multi-stream-grpc
docker build -t grpc .
```

To run the container, we need to map the `$GOPATH` directory on the host into the `/go` directory in the container, because the Dockerfile configures `/go` as the `$GOPATH` in the container:

```bash
docker run --rm -it -v $GOPATH:/go -w /go/src/github.com/wangkuiyi/multi-stream-grpc grpc bash
```

Now, in the container, we can compile the `sqlflow.proto` in this repo into the Go source code:

```bash
protoc sqlflow.proto --go_out=plugins=grpc:.
```

Similarly, we can compile it into Python:

```bash
python -m grpc_tools.protoc -I . --python_out=. --grpc_python_out=. sqlflow.proto
```
