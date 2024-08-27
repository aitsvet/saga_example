// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: v1/mp.proto

package mp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Marketplace_Deposit_FullMethodName    = "/marketplace.Marketplace/Deposit"
	Marketplace_Balance_FullMethodName    = "/marketplace.Marketplace/Balance"
	Marketplace_CreateAd_FullMethodName   = "/marketplace.Marketplace/CreateAd"
	Marketplace_ListAds_FullMethodName    = "/marketplace.Marketplace/ListAds"
	Marketplace_Order_FullMethodName      = "/marketplace.Marketplace/Order"
	Marketplace_ListOrders_FullMethodName = "/marketplace.Marketplace/ListOrders"
	Marketplace_Confirm_FullMethodName    = "/marketplace.Marketplace/Confirm"
	Marketplace_Decline_FullMethodName    = "/marketplace.Marketplace/Decline"
)

// MarketplaceClient is the client API for Marketplace service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketplaceClient interface {
	Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*Empty, error)
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	CreateAd(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*Empty, error)
	ListAds(ctx context.Context, in *ListAdsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[AdResponse], error)
	Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*Empty, error)
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderResponse], error)
	Confirm(ctx context.Context, in *DecisionRequest, opts ...grpc.CallOption) (*Empty, error)
	Decline(ctx context.Context, in *DecisionRequest, opts ...grpc.CallOption) (*Empty, error)
}

type marketplaceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketplaceClient(cc grpc.ClientConnInterface) MarketplaceClient {
	return &marketplaceClient{cc}
}

func (c *marketplaceClient) Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Marketplace_Deposit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceClient) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, Marketplace_Balance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceClient) CreateAd(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Marketplace_CreateAd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceClient) ListAds(ctx context.Context, in *ListAdsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[AdResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Marketplace_ServiceDesc.Streams[0], Marketplace_ListAds_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListAdsRequest, AdResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Marketplace_ListAdsClient = grpc.ServerStreamingClient[AdResponse]

func (c *marketplaceClient) Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Marketplace_Order_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Marketplace_ServiceDesc.Streams[1], Marketplace_ListOrders_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListOrdersRequest, OrderResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Marketplace_ListOrdersClient = grpc.ServerStreamingClient[OrderResponse]

func (c *marketplaceClient) Confirm(ctx context.Context, in *DecisionRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Marketplace_Confirm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceClient) Decline(ctx context.Context, in *DecisionRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Marketplace_Decline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketplaceServer is the server API for Marketplace service.
// All implementations must embed UnimplementedMarketplaceServer
// for forward compatibility.
type MarketplaceServer interface {
	Deposit(context.Context, *DepositRequest) (*Empty, error)
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	CreateAd(context.Context, *AdRequest) (*Empty, error)
	ListAds(*ListAdsRequest, grpc.ServerStreamingServer[AdResponse]) error
	Order(context.Context, *OrderRequest) (*Empty, error)
	ListOrders(*ListOrdersRequest, grpc.ServerStreamingServer[OrderResponse]) error
	Confirm(context.Context, *DecisionRequest) (*Empty, error)
	Decline(context.Context, *DecisionRequest) (*Empty, error)
	mustEmbedUnimplementedMarketplaceServer()
}

// UnimplementedMarketplaceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMarketplaceServer struct{}

func (UnimplementedMarketplaceServer) Deposit(context.Context, *DepositRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedMarketplaceServer) Balance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedMarketplaceServer) CreateAd(context.Context, *AdRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAd not implemented")
}
func (UnimplementedMarketplaceServer) ListAds(*ListAdsRequest, grpc.ServerStreamingServer[AdResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ListAds not implemented")
}
func (UnimplementedMarketplaceServer) Order(context.Context, *OrderRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Order not implemented")
}
func (UnimplementedMarketplaceServer) ListOrders(*ListOrdersRequest, grpc.ServerStreamingServer[OrderResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedMarketplaceServer) Confirm(context.Context, *DecisionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirm not implemented")
}
func (UnimplementedMarketplaceServer) Decline(context.Context, *DecisionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decline not implemented")
}
func (UnimplementedMarketplaceServer) mustEmbedUnimplementedMarketplaceServer() {}
func (UnimplementedMarketplaceServer) testEmbeddedByValue()                     {}

// UnsafeMarketplaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketplaceServer will
// result in compilation errors.
type UnsafeMarketplaceServer interface {
	mustEmbedUnimplementedMarketplaceServer()
}

func RegisterMarketplaceServer(s grpc.ServiceRegistrar, srv MarketplaceServer) {
	// If the following call pancis, it indicates UnimplementedMarketplaceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Marketplace_ServiceDesc, srv)
}

func _Marketplace_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Marketplace_Deposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).Deposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Marketplace_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Marketplace_Balance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Marketplace_CreateAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).CreateAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Marketplace_CreateAd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).CreateAd(ctx, req.(*AdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Marketplace_ListAds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAdsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketplaceServer).ListAds(m, &grpc.GenericServerStream[ListAdsRequest, AdResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Marketplace_ListAdsServer = grpc.ServerStreamingServer[AdResponse]

func _Marketplace_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Marketplace_Order_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).Order(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Marketplace_ListOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListOrdersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketplaceServer).ListOrders(m, &grpc.GenericServerStream[ListOrdersRequest, OrderResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Marketplace_ListOrdersServer = grpc.ServerStreamingServer[OrderResponse]

func _Marketplace_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Marketplace_Confirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).Confirm(ctx, req.(*DecisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Marketplace_Decline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).Decline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Marketplace_Decline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).Decline(ctx, req.(*DecisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Marketplace_ServiceDesc is the grpc.ServiceDesc for Marketplace service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Marketplace_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "marketplace.Marketplace",
	HandlerType: (*MarketplaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deposit",
			Handler:    _Marketplace_Deposit_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _Marketplace_Balance_Handler,
		},
		{
			MethodName: "CreateAd",
			Handler:    _Marketplace_CreateAd_Handler,
		},
		{
			MethodName: "Order",
			Handler:    _Marketplace_Order_Handler,
		},
		{
			MethodName: "Confirm",
			Handler:    _Marketplace_Confirm_Handler,
		},
		{
			MethodName: "Decline",
			Handler:    _Marketplace_Decline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAds",
			Handler:       _Marketplace_ListAds_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListOrders",
			Handler:       _Marketplace_ListOrders_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/mp.proto",
}
