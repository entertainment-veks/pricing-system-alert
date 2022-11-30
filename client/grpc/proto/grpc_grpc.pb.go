// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/grpc.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PricingAlertClient is the client API for PricingAlert service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PricingAlertClient interface {
	SubscribeOnAlerts(ctx context.Context, in *SubscribeOnAlertsParams, opts ...grpc.CallOption) (PricingAlert_SubscribeOnAlertsClient, error)
}

type pricingAlertClient struct {
	cc grpc.ClientConnInterface
}

func NewPricingAlertClient(cc grpc.ClientConnInterface) PricingAlertClient {
	return &pricingAlertClient{cc}
}

func (c *pricingAlertClient) SubscribeOnAlerts(ctx context.Context, in *SubscribeOnAlertsParams, opts ...grpc.CallOption) (PricingAlert_SubscribeOnAlertsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PricingAlert_ServiceDesc.Streams[0], "/grpc.PricingAlert/SubscribeOnAlerts", opts...)
	if err != nil {
		return nil, err
	}
	x := &pricingAlertSubscribeOnAlertsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PricingAlert_SubscribeOnAlertsClient interface {
	Recv() (*PriceNote, error)
	grpc.ClientStream
}

type pricingAlertSubscribeOnAlertsClient struct {
	grpc.ClientStream
}

func (x *pricingAlertSubscribeOnAlertsClient) Recv() (*PriceNote, error) {
	m := new(PriceNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PricingAlertServer is the server API for PricingAlert service.
// All implementations must embed UnimplementedPricingAlertServer
// for forward compatibility
type PricingAlertServer interface {
	SubscribeOnAlerts(*SubscribeOnAlertsParams, PricingAlert_SubscribeOnAlertsServer) error
	mustEmbedUnimplementedPricingAlertServer()
}

// UnimplementedPricingAlertServer must be embedded to have forward compatible implementations.
type UnimplementedPricingAlertServer struct {
}

func (UnimplementedPricingAlertServer) SubscribeOnAlerts(*SubscribeOnAlertsParams, PricingAlert_SubscribeOnAlertsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeOnAlerts not implemented")
}
func (UnimplementedPricingAlertServer) mustEmbedUnimplementedPricingAlertServer() {}

// UnsafePricingAlertServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PricingAlertServer will
// result in compilation errors.
type UnsafePricingAlertServer interface {
	mustEmbedUnimplementedPricingAlertServer()
}

func RegisterPricingAlertServer(s grpc.ServiceRegistrar, srv PricingAlertServer) {
	s.RegisterService(&PricingAlert_ServiceDesc, srv)
}

func _PricingAlert_SubscribeOnAlerts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeOnAlertsParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PricingAlertServer).SubscribeOnAlerts(m, &pricingAlertSubscribeOnAlertsServer{stream})
}

type PricingAlert_SubscribeOnAlertsServer interface {
	Send(*PriceNote) error
	grpc.ServerStream
}

type pricingAlertSubscribeOnAlertsServer struct {
	grpc.ServerStream
}

func (x *pricingAlertSubscribeOnAlertsServer) Send(m *PriceNote) error {
	return x.ServerStream.SendMsg(m)
}

// PricingAlert_ServiceDesc is the grpc.ServiceDesc for PricingAlert service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PricingAlert_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.PricingAlert",
	HandlerType: (*PricingAlertServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeOnAlerts",
			Handler:       _PricingAlert_SubscribeOnAlerts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/grpc.proto",
}
