// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: vesseloracle/vesseloracle/tx.proto

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
	Msg_UpdateParams_FullMethodName                 = "/vesseloracle.vesseloracle.Msg/UpdateParams"
	Msg_CreateVessel_FullMethodName                 = "/vesseloracle.vesseloracle.Msg/CreateVessel"
	Msg_UpdateVessel_FullMethodName                 = "/vesseloracle.vesseloracle.Msg/UpdateVessel"
	Msg_DeleteVessel_FullMethodName                 = "/vesseloracle.vesseloracle.Msg/DeleteVessel"
	Msg_ConsolidateReports_FullMethodName           = "/vesseloracle.vesseloracle.Msg/ConsolidateReports"
	Msg_CreateConsolidatedDataReport_FullMethodName = "/vesseloracle.vesseloracle.Msg/CreateConsolidatedDataReport"
	Msg_UpdateConsolidatedDataReport_FullMethodName = "/vesseloracle.vesseloracle.Msg/UpdateConsolidatedDataReport"
	Msg_DeleteConsolidatedDataReport_FullMethodName = "/vesseloracle.vesseloracle.Msg/DeleteConsolidatedDataReport"
	Msg_TransmitReport_FullMethodName               = "/vesseloracle.vesseloracle.Msg/TransmitReport"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateVessel(ctx context.Context, in *MsgCreateVessel, opts ...grpc.CallOption) (*MsgCreateVesselResponse, error)
	UpdateVessel(ctx context.Context, in *MsgUpdateVessel, opts ...grpc.CallOption) (*MsgUpdateVesselResponse, error)
	DeleteVessel(ctx context.Context, in *MsgDeleteVessel, opts ...grpc.CallOption) (*MsgDeleteVesselResponse, error)
	ConsolidateReports(ctx context.Context, in *MsgConsolidateReports, opts ...grpc.CallOption) (*MsgConsolidateReportsResponse, error)
	CreateConsolidatedDataReport(ctx context.Context, in *MsgCreateConsolidatedDataReport, opts ...grpc.CallOption) (*MsgCreateConsolidatedDataReportResponse, error)
	UpdateConsolidatedDataReport(ctx context.Context, in *MsgUpdateConsolidatedDataReport, opts ...grpc.CallOption) (*MsgUpdateConsolidatedDataReportResponse, error)
	DeleteConsolidatedDataReport(ctx context.Context, in *MsgDeleteConsolidatedDataReport, opts ...grpc.CallOption) (*MsgDeleteConsolidatedDataReportResponse, error)
	TransmitReport(ctx context.Context, in *MsgTransmitReport, opts ...grpc.CallOption) (*MsgTransmitReportResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateVessel(ctx context.Context, in *MsgCreateVessel, opts ...grpc.CallOption) (*MsgCreateVesselResponse, error) {
	out := new(MsgCreateVesselResponse)
	err := c.cc.Invoke(ctx, Msg_CreateVessel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateVessel(ctx context.Context, in *MsgUpdateVessel, opts ...grpc.CallOption) (*MsgUpdateVesselResponse, error) {
	out := new(MsgUpdateVesselResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateVessel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteVessel(ctx context.Context, in *MsgDeleteVessel, opts ...grpc.CallOption) (*MsgDeleteVesselResponse, error) {
	out := new(MsgDeleteVesselResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteVessel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConsolidateReports(ctx context.Context, in *MsgConsolidateReports, opts ...grpc.CallOption) (*MsgConsolidateReportsResponse, error) {
	out := new(MsgConsolidateReportsResponse)
	err := c.cc.Invoke(ctx, Msg_ConsolidateReports_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateConsolidatedDataReport(ctx context.Context, in *MsgCreateConsolidatedDataReport, opts ...grpc.CallOption) (*MsgCreateConsolidatedDataReportResponse, error) {
	out := new(MsgCreateConsolidatedDataReportResponse)
	err := c.cc.Invoke(ctx, Msg_CreateConsolidatedDataReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateConsolidatedDataReport(ctx context.Context, in *MsgUpdateConsolidatedDataReport, opts ...grpc.CallOption) (*MsgUpdateConsolidatedDataReportResponse, error) {
	out := new(MsgUpdateConsolidatedDataReportResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateConsolidatedDataReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteConsolidatedDataReport(ctx context.Context, in *MsgDeleteConsolidatedDataReport, opts ...grpc.CallOption) (*MsgDeleteConsolidatedDataReportResponse, error) {
	out := new(MsgDeleteConsolidatedDataReportResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteConsolidatedDataReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransmitReport(ctx context.Context, in *MsgTransmitReport, opts ...grpc.CallOption) (*MsgTransmitReportResponse, error) {
	out := new(MsgTransmitReportResponse)
	err := c.cc.Invoke(ctx, Msg_TransmitReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateVessel(context.Context, *MsgCreateVessel) (*MsgCreateVesselResponse, error)
	UpdateVessel(context.Context, *MsgUpdateVessel) (*MsgUpdateVesselResponse, error)
	DeleteVessel(context.Context, *MsgDeleteVessel) (*MsgDeleteVesselResponse, error)
	ConsolidateReports(context.Context, *MsgConsolidateReports) (*MsgConsolidateReportsResponse, error)
	CreateConsolidatedDataReport(context.Context, *MsgCreateConsolidatedDataReport) (*MsgCreateConsolidatedDataReportResponse, error)
	UpdateConsolidatedDataReport(context.Context, *MsgUpdateConsolidatedDataReport) (*MsgUpdateConsolidatedDataReportResponse, error)
	DeleteConsolidatedDataReport(context.Context, *MsgDeleteConsolidatedDataReport) (*MsgDeleteConsolidatedDataReportResponse, error)
	TransmitReport(context.Context, *MsgTransmitReport) (*MsgTransmitReportResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateVessel(context.Context, *MsgCreateVessel) (*MsgCreateVesselResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVessel not implemented")
}
func (UnimplementedMsgServer) UpdateVessel(context.Context, *MsgUpdateVessel) (*MsgUpdateVesselResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVessel not implemented")
}
func (UnimplementedMsgServer) DeleteVessel(context.Context, *MsgDeleteVessel) (*MsgDeleteVesselResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVessel not implemented")
}
func (UnimplementedMsgServer) ConsolidateReports(context.Context, *MsgConsolidateReports) (*MsgConsolidateReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsolidateReports not implemented")
}
func (UnimplementedMsgServer) CreateConsolidatedDataReport(context.Context, *MsgCreateConsolidatedDataReport) (*MsgCreateConsolidatedDataReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsolidatedDataReport not implemented")
}
func (UnimplementedMsgServer) UpdateConsolidatedDataReport(context.Context, *MsgUpdateConsolidatedDataReport) (*MsgUpdateConsolidatedDataReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConsolidatedDataReport not implemented")
}
func (UnimplementedMsgServer) DeleteConsolidatedDataReport(context.Context, *MsgDeleteConsolidatedDataReport) (*MsgDeleteConsolidatedDataReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConsolidatedDataReport not implemented")
}
func (UnimplementedMsgServer) TransmitReport(context.Context, *MsgTransmitReport) (*MsgTransmitReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransmitReport not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateVessel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateVessel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateVessel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateVessel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateVessel(ctx, req.(*MsgCreateVessel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateVessel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateVessel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateVessel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateVessel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateVessel(ctx, req.(*MsgUpdateVessel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteVessel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteVessel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteVessel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteVessel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteVessel(ctx, req.(*MsgDeleteVessel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConsolidateReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConsolidateReports)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConsolidateReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ConsolidateReports_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConsolidateReports(ctx, req.(*MsgConsolidateReports))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateConsolidatedDataReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateConsolidatedDataReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateConsolidatedDataReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateConsolidatedDataReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateConsolidatedDataReport(ctx, req.(*MsgCreateConsolidatedDataReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateConsolidatedDataReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateConsolidatedDataReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateConsolidatedDataReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateConsolidatedDataReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateConsolidatedDataReport(ctx, req.(*MsgUpdateConsolidatedDataReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteConsolidatedDataReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteConsolidatedDataReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteConsolidatedDataReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteConsolidatedDataReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteConsolidatedDataReport(ctx, req.(*MsgDeleteConsolidatedDataReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransmitReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransmitReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransmitReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransmitReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransmitReport(ctx, req.(*MsgTransmitReport))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vesseloracle.vesseloracle.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateVessel",
			Handler:    _Msg_CreateVessel_Handler,
		},
		{
			MethodName: "UpdateVessel",
			Handler:    _Msg_UpdateVessel_Handler,
		},
		{
			MethodName: "DeleteVessel",
			Handler:    _Msg_DeleteVessel_Handler,
		},
		{
			MethodName: "ConsolidateReports",
			Handler:    _Msg_ConsolidateReports_Handler,
		},
		{
			MethodName: "CreateConsolidatedDataReport",
			Handler:    _Msg_CreateConsolidatedDataReport_Handler,
		},
		{
			MethodName: "UpdateConsolidatedDataReport",
			Handler:    _Msg_UpdateConsolidatedDataReport_Handler,
		},
		{
			MethodName: "DeleteConsolidatedDataReport",
			Handler:    _Msg_DeleteConsolidatedDataReport_Handler,
		},
		{
			MethodName: "TransmitReport",
			Handler:    _Msg_TransmitReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vesseloracle/vesseloracle/tx.proto",
}
