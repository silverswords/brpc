package mast

import "github.com/silverswords/mast/mastgrpc"

// Builder could build for given parameters to make
// rpc and grpc client and server
type Builder interface {
	Client() Client
	Server() Server
}

type Builders struct {
	*mastgrpc.GrpcBuilder
}

// Options configure Builder's options
type Options func(string) error

// Server could listen and serve
type Server interface {
	Prepare(info, registerFunc interface{})
	Start()
}

// Client supposed Synchronous and Asynchronous
type Client interface {
	Call()
	Go()
}

// BuildClient return Client
func BuildClient(b Builder) Client {
	return b.Client()
}

// BuildServer return server
func BuildServer(b Builder) Server {
	return b.Server()
}

// Config configure options in rpc
func Config(opt string) {

}
