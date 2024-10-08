// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: vesseloracle/vesseloracle/query.proto

package vesseloracle

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

const (
	Query_Params_FullMethodName                    = "/vesseloracle.vesseloracle.Query/Params"
	Query_Vessel_FullMethodName                    = "/vesseloracle.vesseloracle.Query/Vessel"
	Query_VesselAll_FullMethodName                 = "/vesseloracle.vesseloracle.Query/VesselAll"
	Query_ConsolidatedDataReport_FullMethodName    = "/vesseloracle.vesseloracle.Query/ConsolidatedDataReport"
	Query_ConsolidatedDataReportAll_FullMethodName = "/vesseloracle.vesseloracle.Query/ConsolidatedDataReportAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Vessel items.
	Vessel(ctx context.Context, in *QueryGetVesselRequest, opts ...grpc.CallOption) (*QueryGetVesselResponse, error)
	VesselAll(ctx context.Context, in *QueryAllVesselRequest, opts ...grpc.CallOption) (*QueryAllVesselResponse, error)
	// Queries a list of ConsolidatedDataReport items.
	ConsolidatedDataReport(ctx context.Context, in *QueryGetConsolidatedDataReportRequest, opts ...grpc.CallOption) (*QueryGetConsolidatedDataReportResponse, error)
	ConsolidatedDataReportAll(ctx context.Context, in *QueryAllConsolidatedDataReportRequest, opts ...grpc.CallOption) (*QueryAllConsolidatedDataReportResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Vessel(ctx context.Context, in *QueryGetVesselRequest, opts ...grpc.CallOption) (*QueryGetVesselResponse, error) {
	out := new(QueryGetVesselResponse)
	err := c.cc.Invoke(ctx, Query_Vessel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VesselAll(ctx context.Context, in *QueryAllVesselRequest, opts ...grpc.CallOption) (*QueryAllVesselResponse, error) {
	out := new(QueryAllVesselResponse)
	err := c.cc.Invoke(ctx, Query_VesselAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ConsolidatedDataReport(ctx context.Context, in *QueryGetConsolidatedDataReportRequest, opts ...grpc.CallOption) (*QueryGetConsolidatedDataReportResponse, error) {
	out := new(QueryGetConsolidatedDataReportResponse)
	err := c.cc.Invoke(ctx, Query_ConsolidatedDataReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ConsolidatedDataReportAll(ctx context.Context, in *QueryAllConsolidatedDataReportRequest, opts ...grpc.CallOption) (*QueryAllConsolidatedDataReportResponse, error) {
	out := new(QueryAllConsolidatedDataReportResponse)
	err := c.cc.Invoke(ctx, Query_ConsolidatedDataReportAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Vessel items.
	Vessel(context.Context, *QueryGetVesselRequest) (*QueryGetVesselResponse, error)
	VesselAll(context.Context, *QueryAllVesselRequest) (*QueryAllVesselResponse, error)
	// Queries a list of ConsolidatedDataReport items.
	ConsolidatedDataReport(context.Context, *QueryGetConsolidatedDataReportRequest) (*QueryGetConsolidatedDataReportResponse, error)
	ConsolidatedDataReportAll(context.Context, *QueryAllConsolidatedDataReportRequest) (*QueryAllConsolidatedDataReportResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Vessel(context.Context, *QueryGetVesselRequest) (*QueryGetVesselResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vessel not implemented")
}
func (UnimplementedQueryServer) VesselAll(context.Context, *QueryAllVesselRequest) (*QueryAllVesselResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VesselAll not implemented")
}
func (UnimplementedQueryServer) ConsolidatedDataReport(context.Context, *QueryGetConsolidatedDataReportRequest) (*QueryGetConsolidatedDataReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsolidatedDataReport not implemented")
}
func (UnimplementedQueryServer) ConsolidatedDataReportAll(context.Context, *QueryAllConsolidatedDataReportRequest) (*QueryAllConsolidatedDataReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsolidatedDataReportAll not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Vessel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetVesselRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Vessel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Vessel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Vessel(ctx, req.(*QueryGetVesselRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VesselAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllVesselRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VesselAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_VesselAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VesselAll(ctx, req.(*QueryAllVesselRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ConsolidatedDataReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetConsolidatedDataReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ConsolidatedDataReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ConsolidatedDataReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ConsolidatedDataReport(ctx, req.(*QueryGetConsolidatedDataReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ConsolidatedDataReportAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllConsolidatedDataReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ConsolidatedDataReportAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ConsolidatedDataReportAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ConsolidatedDataReportAll(ctx, req.(*QueryAllConsolidatedDataReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vesseloracle.vesseloracle.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Vessel",
			Handler:    _Query_Vessel_Handler,
		},
		{
			MethodName: "VesselAll",
			Handler:    _Query_VesselAll_Handler,
		},
		{
			MethodName: "ConsolidatedDataReport",
			Handler:    _Query_ConsolidatedDataReport_Handler,
		},
		{
			MethodName: "ConsolidatedDataReportAll",
			Handler:    _Query_ConsolidatedDataReportAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vesseloracle/vesseloracle/query.proto",
}
