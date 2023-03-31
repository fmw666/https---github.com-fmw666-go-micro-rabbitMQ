// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderService.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for OrderService service

func NewOrderServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderService service

type OrderService interface {
	GetOrderList(ctx context.Context, in *OrderListRequest, opts ...client.CallOption) (*OrderListResponse, error)
	CreateOrder(ctx context.Context, in *OrderCreateRequest, opts ...client.CallOption) (*OrderCreateResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) GetOrderList(ctx context.Context, in *OrderListRequest, opts ...client.CallOption) (*OrderListResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetOrderList", in)
	out := new(OrderListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CreateOrder(ctx context.Context, in *OrderCreateRequest, opts ...client.CallOption) (*OrderCreateResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.CreateOrder", in)
	out := new(OrderCreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	GetOrderList(context.Context, *OrderListRequest, *OrderListResponse) error
	CreateOrder(context.Context, *OrderCreateRequest, *OrderCreateResponse) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) error {
	type orderService interface {
		GetOrderList(ctx context.Context, in *OrderListRequest, out *OrderListResponse) error
		CreateOrder(ctx context.Context, in *OrderCreateRequest, out *OrderCreateResponse) error
	}
	type OrderService struct {
		orderService
	}
	h := &orderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderService{h}, opts...))
}

type orderServiceHandler struct {
	OrderServiceHandler
}

func (h *orderServiceHandler) GetOrderList(ctx context.Context, in *OrderListRequest, out *OrderListResponse) error {
	return h.OrderServiceHandler.GetOrderList(ctx, in, out)
}

func (h *orderServiceHandler) CreateOrder(ctx context.Context, in *OrderCreateRequest, out *OrderCreateResponse) error {
	return h.OrderServiceHandler.CreateOrder(ctx, in, out)
}