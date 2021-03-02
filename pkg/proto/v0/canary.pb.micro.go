// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: canary.proto

package proto

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Canary service

func NewCanaryEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "Canary.SetCanary",
			Path:    []string{"/api/v0/canary"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for Canary service

type CanaryService interface {
	SetCanary(ctx context.Context, in *CanaryRequest, opts ...client.CallOption) (*CanaryResponse, error)
}

type canaryService struct {
	c    client.Client
	name string
}

func NewCanaryService(name string, c client.Client) CanaryService {
	return &canaryService{
		c:    c,
		name: name,
	}
}

func (c *canaryService) SetCanary(ctx context.Context, in *CanaryRequest, opts ...client.CallOption) (*CanaryResponse, error) {
	req := c.c.NewRequest(c.name, "Canary.SetCanary", in)
	out := new(CanaryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Canary service

type CanaryHandler interface {
	SetCanary(context.Context, *CanaryRequest, *CanaryResponse) error
}

func RegisterCanaryHandler(s server.Server, hdlr CanaryHandler, opts ...server.HandlerOption) error {
	type canary interface {
		SetCanary(ctx context.Context, in *CanaryRequest, out *CanaryResponse) error
	}
	type Canary struct {
		canary
	}
	h := &canaryHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Canary.SetCanary",
		Path:    []string{"/api/v0/canary"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&Canary{h}, opts...))
}

type canaryHandler struct {
	CanaryHandler
}

func (h *canaryHandler) SetCanary(ctx context.Context, in *CanaryRequest, out *CanaryResponse) error {
	return h.CanaryHandler.SetCanary(ctx, in, out)
}
