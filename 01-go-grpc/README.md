# A simple gRPC application in Go
This application implements a send message service that allows the client to invoke a method on the server i.e. SendMessage. The server recives the message and replies back to the client with an acknowledgement. The intent of the application is to get familiar with gRPC using Go.

# Pre-requisites
gRPC uses **Protocol Buffers** and hence, you will first need to install it. On Mac OSX, you can install it using HomeBrew:
```bash
$ brew install protobuf
```
Alternatively, you can follow this [guide](http://google.github.io/proto-lens/installing-protoc.html) to get the `protoc` cli

You'll then need the protocol compiler plugin for Go
```bash
$ go get github.com/golang/protobuf/protoc-gen-go
```
Update your `$PATH` so that `protoc` can find the plugin
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```
Get the `grpc` library
```bash
$ go get google.golang.org/grpc
```
# Generating the gRPC code
```bash
$ protoc --go_out=plugins=grpc:chat chat.proto
```

# Running the application
Open a terminal and run the server
```bash
$ go run server/server.go
```
Then, in another terminal, execute the client by passing a message as an argument
```bash
$ go run client/client.go "Hello."
```