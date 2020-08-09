// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: model/proto/notification/notification.proto

package legoas_srv_notification

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

// Api Endpoints for Notification service

func NewNotificationEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Notification service

type NotificationService interface {
	Send(ctx context.Context, in *SendRequest, opts ...client.CallOption) (*SendResponse, error)
}

type notificationService struct {
	c    client.Client
	name string
}

func NewNotificationService(name string, c client.Client) NotificationService {
	return &notificationService{
		c:    c,
		name: name,
	}
}

func (c *notificationService) Send(ctx context.Context, in *SendRequest, opts ...client.CallOption) (*SendResponse, error) {
	req := c.c.NewRequest(c.name, "Notification.Send", in)
	out := new(SendResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notification service

type NotificationHandler interface {
	Send(context.Context, *SendRequest, *SendResponse) error
}

func RegisterNotificationHandler(s server.Server, hdlr NotificationHandler, opts ...server.HandlerOption) error {
	type notification interface {
		Send(ctx context.Context, in *SendRequest, out *SendResponse) error
	}
	type Notification struct {
		notification
	}
	h := &notificationHandler{hdlr}
	return s.Handle(s.NewHandler(&Notification{h}, opts...))
}

type notificationHandler struct {
	NotificationHandler
}

func (h *notificationHandler) Send(ctx context.Context, in *SendRequest, out *SendResponse) error {
	return h.NotificationHandler.Send(ctx, in, out)
}