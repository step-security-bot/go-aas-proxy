// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gen/proto/basyx_openapi.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BasyxOpenapiClient is the client API for BasyxOpenapi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BasyxOpenapiClient interface {
	GetAssetAdministrationShell(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AssetAdministrationShellDescriptor, error)
	GetSubmodelsFromShell(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Submodel, error)
	GETAasSubmodelsSubmodelIdShort(ctx context.Context, in *GETAasSubmodelsSubmodelIdShortParameters, opts ...grpc.CallOption) (*Submodel, error)
	PutSubmodelToShell(ctx context.Context, in *PutSubmodelToShellParameters, opts ...grpc.CallOption) (*Submodel, error)
	DeleteSubmodelFromShellByIdShort(ctx context.Context, in *DeleteSubmodelFromShellByIdShortParameters, opts ...grpc.CallOption) (*Result, error)
	GetSubmodelFromShellByIdShort(ctx context.Context, in *GetSubmodelFromShellByIdShortParameters, opts ...grpc.CallOption) (*Submodel, error)
	ShellGetSubmodelValues(ctx context.Context, in *ShellGetSubmodelValuesParameters, opts ...grpc.CallOption) (*Result, error)
	ShellGetSubmodelElements(ctx context.Context, in *ShellGetSubmodelElementsParameters, opts ...grpc.CallOption) (*SubmodelElement, error)
	ShellGetSubmodelElementByIdShort(ctx context.Context, in *ShellGetSubmodelElementByIdShortParameters, opts ...grpc.CallOption) (*SubmodelElement, error)
	ShellPutSubmodelElement(ctx context.Context, in *ShellPutSubmodelElementParameters, opts ...grpc.CallOption) (*SubmodelElement, error)
	ShellDeleteSubmodelElementByIdShort(ctx context.Context, in *ShellDeleteSubmodelElementByIdShortParameters, opts ...grpc.CallOption) (*Result, error)
	ShellGetSubmodelElementValueByIdShort(ctx context.Context, in *ShellGetSubmodelElementValueByIdShortParameters, opts ...grpc.CallOption) (*ShellGetSubmodelElementValueByIdShortOK, error)
	ShellPutSubmodelElementValueByIdShort(ctx context.Context, in *ShellPutSubmodelElementValueByIdShortParameters, opts ...grpc.CallOption) (*ElementValue, error)
	ShellInvokeOperationByIdShort(ctx context.Context, in *ShellInvokeOperationByIdShortParameters, opts ...grpc.CallOption) (*Result, error)
	ShellGetInvocationResultByIdShort(ctx context.Context, in *ShellGetInvocationResultByIdShortParameters, opts ...grpc.CallOption) (*InvocationResponse, error)
}

type basyxOpenapiClient struct {
	cc grpc.ClientConnInterface
}

func NewBasyxOpenapiClient(cc grpc.ClientConnInterface) BasyxOpenapiClient {
	return &basyxOpenapiClient{cc}
}

func (c *basyxOpenapiClient) GetAssetAdministrationShell(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AssetAdministrationShellDescriptor, error) {
	out := new(AssetAdministrationShellDescriptor)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/GetAssetAdministrationShell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) GetSubmodelsFromShell(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Submodel, error) {
	out := new(Submodel)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/GetSubmodelsFromShell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) GETAasSubmodelsSubmodelIdShort(ctx context.Context, in *GETAasSubmodelsSubmodelIdShortParameters, opts ...grpc.CallOption) (*Submodel, error) {
	out := new(Submodel)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/GETAasSubmodelsSubmodelIdShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) PutSubmodelToShell(ctx context.Context, in *PutSubmodelToShellParameters, opts ...grpc.CallOption) (*Submodel, error) {
	out := new(Submodel)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/PutSubmodelToShell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) DeleteSubmodelFromShellByIdShort(ctx context.Context, in *DeleteSubmodelFromShellByIdShortParameters, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/DeleteSubmodelFromShellByIdShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) GetSubmodelFromShellByIdShort(ctx context.Context, in *GetSubmodelFromShellByIdShortParameters, opts ...grpc.CallOption) (*Submodel, error) {
	out := new(Submodel)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/GetSubmodelFromShellByIdShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) ShellGetSubmodelValues(ctx context.Context, in *ShellGetSubmodelValuesParameters, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/ShellGetSubmodelValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) ShellGetSubmodelElements(ctx context.Context, in *ShellGetSubmodelElementsParameters, opts ...grpc.CallOption) (*SubmodelElement, error) {
	out := new(SubmodelElement)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/ShellGetSubmodelElements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) ShellGetSubmodelElementByIdShort(ctx context.Context, in *ShellGetSubmodelElementByIdShortParameters, opts ...grpc.CallOption) (*SubmodelElement, error) {
	out := new(SubmodelElement)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/ShellGetSubmodelElementByIdShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) ShellPutSubmodelElement(ctx context.Context, in *ShellPutSubmodelElementParameters, opts ...grpc.CallOption) (*SubmodelElement, error) {
	out := new(SubmodelElement)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/ShellPutSubmodelElement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) ShellDeleteSubmodelElementByIdShort(ctx context.Context, in *ShellDeleteSubmodelElementByIdShortParameters, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/ShellDeleteSubmodelElementByIdShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) ShellGetSubmodelElementValueByIdShort(ctx context.Context, in *ShellGetSubmodelElementValueByIdShortParameters, opts ...grpc.CallOption) (*ShellGetSubmodelElementValueByIdShortOK, error) {
	out := new(ShellGetSubmodelElementValueByIdShortOK)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/ShellGetSubmodelElementValueByIdShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) ShellPutSubmodelElementValueByIdShort(ctx context.Context, in *ShellPutSubmodelElementValueByIdShortParameters, opts ...grpc.CallOption) (*ElementValue, error) {
	out := new(ElementValue)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/ShellPutSubmodelElementValueByIdShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) ShellInvokeOperationByIdShort(ctx context.Context, in *ShellInvokeOperationByIdShortParameters, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/ShellInvokeOperationByIdShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basyxOpenapiClient) ShellGetInvocationResultByIdShort(ctx context.Context, in *ShellGetInvocationResultByIdShortParameters, opts ...grpc.CallOption) (*InvocationResponse, error) {
	out := new(InvocationResponse)
	err := c.cc.Invoke(ctx, "/basyx_openapi.Basyx_openapi/ShellGetInvocationResultByIdShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasyxOpenapiServer is the server API for BasyxOpenapi service.
// All implementations should embed UnimplementedBasyxOpenapiServer
// for forward compatibility
type BasyxOpenapiServer interface {
	GetAssetAdministrationShell(context.Context, *emptypb.Empty) (*AssetAdministrationShellDescriptor, error)
	GetSubmodelsFromShell(context.Context, *emptypb.Empty) (*Submodel, error)
	GETAasSubmodelsSubmodelIdShort(context.Context, *GETAasSubmodelsSubmodelIdShortParameters) (*Submodel, error)
	PutSubmodelToShell(context.Context, *PutSubmodelToShellParameters) (*Submodel, error)
	DeleteSubmodelFromShellByIdShort(context.Context, *DeleteSubmodelFromShellByIdShortParameters) (*Result, error)
	GetSubmodelFromShellByIdShort(context.Context, *GetSubmodelFromShellByIdShortParameters) (*Submodel, error)
	ShellGetSubmodelValues(context.Context, *ShellGetSubmodelValuesParameters) (*Result, error)
	ShellGetSubmodelElements(context.Context, *ShellGetSubmodelElementsParameters) (*SubmodelElement, error)
	ShellGetSubmodelElementByIdShort(context.Context, *ShellGetSubmodelElementByIdShortParameters) (*SubmodelElement, error)
	ShellPutSubmodelElement(context.Context, *ShellPutSubmodelElementParameters) (*SubmodelElement, error)
	ShellDeleteSubmodelElementByIdShort(context.Context, *ShellDeleteSubmodelElementByIdShortParameters) (*Result, error)
	ShellGetSubmodelElementValueByIdShort(context.Context, *ShellGetSubmodelElementValueByIdShortParameters) (*ShellGetSubmodelElementValueByIdShortOK, error)
	ShellPutSubmodelElementValueByIdShort(context.Context, *ShellPutSubmodelElementValueByIdShortParameters) (*ElementValue, error)
	ShellInvokeOperationByIdShort(context.Context, *ShellInvokeOperationByIdShortParameters) (*Result, error)
	ShellGetInvocationResultByIdShort(context.Context, *ShellGetInvocationResultByIdShortParameters) (*InvocationResponse, error)
}

// UnimplementedBasyxOpenapiServer should be embedded to have forward compatible implementations.
type UnimplementedBasyxOpenapiServer struct {
}

func (UnimplementedBasyxOpenapiServer) GetAssetAdministrationShell(context.Context, *emptypb.Empty) (*AssetAdministrationShellDescriptor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetAdministrationShell not implemented")
}
func (UnimplementedBasyxOpenapiServer) GetSubmodelsFromShell(context.Context, *emptypb.Empty) (*Submodel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmodelsFromShell not implemented")
}
func (UnimplementedBasyxOpenapiServer) GETAasSubmodelsSubmodelIdShort(context.Context, *GETAasSubmodelsSubmodelIdShortParameters) (*Submodel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GETAasSubmodelsSubmodelIdShort not implemented")
}
func (UnimplementedBasyxOpenapiServer) PutSubmodelToShell(context.Context, *PutSubmodelToShellParameters) (*Submodel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutSubmodelToShell not implemented")
}
func (UnimplementedBasyxOpenapiServer) DeleteSubmodelFromShellByIdShort(context.Context, *DeleteSubmodelFromShellByIdShortParameters) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubmodelFromShellByIdShort not implemented")
}
func (UnimplementedBasyxOpenapiServer) GetSubmodelFromShellByIdShort(context.Context, *GetSubmodelFromShellByIdShortParameters) (*Submodel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmodelFromShellByIdShort not implemented")
}
func (UnimplementedBasyxOpenapiServer) ShellGetSubmodelValues(context.Context, *ShellGetSubmodelValuesParameters) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShellGetSubmodelValues not implemented")
}
func (UnimplementedBasyxOpenapiServer) ShellGetSubmodelElements(context.Context, *ShellGetSubmodelElementsParameters) (*SubmodelElement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShellGetSubmodelElements not implemented")
}
func (UnimplementedBasyxOpenapiServer) ShellGetSubmodelElementByIdShort(context.Context, *ShellGetSubmodelElementByIdShortParameters) (*SubmodelElement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShellGetSubmodelElementByIdShort not implemented")
}
func (UnimplementedBasyxOpenapiServer) ShellPutSubmodelElement(context.Context, *ShellPutSubmodelElementParameters) (*SubmodelElement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShellPutSubmodelElement not implemented")
}
func (UnimplementedBasyxOpenapiServer) ShellDeleteSubmodelElementByIdShort(context.Context, *ShellDeleteSubmodelElementByIdShortParameters) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShellDeleteSubmodelElementByIdShort not implemented")
}
func (UnimplementedBasyxOpenapiServer) ShellGetSubmodelElementValueByIdShort(context.Context, *ShellGetSubmodelElementValueByIdShortParameters) (*ShellGetSubmodelElementValueByIdShortOK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShellGetSubmodelElementValueByIdShort not implemented")
}
func (UnimplementedBasyxOpenapiServer) ShellPutSubmodelElementValueByIdShort(context.Context, *ShellPutSubmodelElementValueByIdShortParameters) (*ElementValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShellPutSubmodelElementValueByIdShort not implemented")
}
func (UnimplementedBasyxOpenapiServer) ShellInvokeOperationByIdShort(context.Context, *ShellInvokeOperationByIdShortParameters) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShellInvokeOperationByIdShort not implemented")
}
func (UnimplementedBasyxOpenapiServer) ShellGetInvocationResultByIdShort(context.Context, *ShellGetInvocationResultByIdShortParameters) (*InvocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShellGetInvocationResultByIdShort not implemented")
}

// UnsafeBasyxOpenapiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BasyxOpenapiServer will
// result in compilation errors.
type UnsafeBasyxOpenapiServer interface {
	mustEmbedUnimplementedBasyxOpenapiServer()
}

func RegisterBasyxOpenapiServer(s grpc.ServiceRegistrar, srv BasyxOpenapiServer) {
	s.RegisterService(&BasyxOpenapi_ServiceDesc, srv)
}

func _BasyxOpenapi_GetAssetAdministrationShell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).GetAssetAdministrationShell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/GetAssetAdministrationShell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).GetAssetAdministrationShell(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_GetSubmodelsFromShell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).GetSubmodelsFromShell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/GetSubmodelsFromShell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).GetSubmodelsFromShell(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_GETAasSubmodelsSubmodelIdShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GETAasSubmodelsSubmodelIdShortParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).GETAasSubmodelsSubmodelIdShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/GETAasSubmodelsSubmodelIdShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).GETAasSubmodelsSubmodelIdShort(ctx, req.(*GETAasSubmodelsSubmodelIdShortParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_PutSubmodelToShell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutSubmodelToShellParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).PutSubmodelToShell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/PutSubmodelToShell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).PutSubmodelToShell(ctx, req.(*PutSubmodelToShellParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_DeleteSubmodelFromShellByIdShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubmodelFromShellByIdShortParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).DeleteSubmodelFromShellByIdShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/DeleteSubmodelFromShellByIdShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).DeleteSubmodelFromShellByIdShort(ctx, req.(*DeleteSubmodelFromShellByIdShortParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_GetSubmodelFromShellByIdShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmodelFromShellByIdShortParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).GetSubmodelFromShellByIdShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/GetSubmodelFromShellByIdShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).GetSubmodelFromShellByIdShort(ctx, req.(*GetSubmodelFromShellByIdShortParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_ShellGetSubmodelValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellGetSubmodelValuesParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).ShellGetSubmodelValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/ShellGetSubmodelValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).ShellGetSubmodelValues(ctx, req.(*ShellGetSubmodelValuesParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_ShellGetSubmodelElements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellGetSubmodelElementsParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).ShellGetSubmodelElements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/ShellGetSubmodelElements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).ShellGetSubmodelElements(ctx, req.(*ShellGetSubmodelElementsParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_ShellGetSubmodelElementByIdShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellGetSubmodelElementByIdShortParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).ShellGetSubmodelElementByIdShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/ShellGetSubmodelElementByIdShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).ShellGetSubmodelElementByIdShort(ctx, req.(*ShellGetSubmodelElementByIdShortParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_ShellPutSubmodelElement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellPutSubmodelElementParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).ShellPutSubmodelElement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/ShellPutSubmodelElement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).ShellPutSubmodelElement(ctx, req.(*ShellPutSubmodelElementParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_ShellDeleteSubmodelElementByIdShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellDeleteSubmodelElementByIdShortParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).ShellDeleteSubmodelElementByIdShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/ShellDeleteSubmodelElementByIdShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).ShellDeleteSubmodelElementByIdShort(ctx, req.(*ShellDeleteSubmodelElementByIdShortParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_ShellGetSubmodelElementValueByIdShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellGetSubmodelElementValueByIdShortParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).ShellGetSubmodelElementValueByIdShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/ShellGetSubmodelElementValueByIdShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).ShellGetSubmodelElementValueByIdShort(ctx, req.(*ShellGetSubmodelElementValueByIdShortParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_ShellPutSubmodelElementValueByIdShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellPutSubmodelElementValueByIdShortParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).ShellPutSubmodelElementValueByIdShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/ShellPutSubmodelElementValueByIdShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).ShellPutSubmodelElementValueByIdShort(ctx, req.(*ShellPutSubmodelElementValueByIdShortParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_ShellInvokeOperationByIdShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellInvokeOperationByIdShortParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).ShellInvokeOperationByIdShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/ShellInvokeOperationByIdShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).ShellInvokeOperationByIdShort(ctx, req.(*ShellInvokeOperationByIdShortParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasyxOpenapi_ShellGetInvocationResultByIdShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellGetInvocationResultByIdShortParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasyxOpenapiServer).ShellGetInvocationResultByIdShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basyx_openapi.Basyx_openapi/ShellGetInvocationResultByIdShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasyxOpenapiServer).ShellGetInvocationResultByIdShort(ctx, req.(*ShellGetInvocationResultByIdShortParameters))
	}
	return interceptor(ctx, in, info, handler)
}

// BasyxOpenapi_ServiceDesc is the grpc.ServiceDesc for BasyxOpenapi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BasyxOpenapi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "basyx_openapi.Basyx_openapi",
	HandlerType: (*BasyxOpenapiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAssetAdministrationShell",
			Handler:    _BasyxOpenapi_GetAssetAdministrationShell_Handler,
		},
		{
			MethodName: "GetSubmodelsFromShell",
			Handler:    _BasyxOpenapi_GetSubmodelsFromShell_Handler,
		},
		{
			MethodName: "GETAasSubmodelsSubmodelIdShort",
			Handler:    _BasyxOpenapi_GETAasSubmodelsSubmodelIdShort_Handler,
		},
		{
			MethodName: "PutSubmodelToShell",
			Handler:    _BasyxOpenapi_PutSubmodelToShell_Handler,
		},
		{
			MethodName: "DeleteSubmodelFromShellByIdShort",
			Handler:    _BasyxOpenapi_DeleteSubmodelFromShellByIdShort_Handler,
		},
		{
			MethodName: "GetSubmodelFromShellByIdShort",
			Handler:    _BasyxOpenapi_GetSubmodelFromShellByIdShort_Handler,
		},
		{
			MethodName: "ShellGetSubmodelValues",
			Handler:    _BasyxOpenapi_ShellGetSubmodelValues_Handler,
		},
		{
			MethodName: "ShellGetSubmodelElements",
			Handler:    _BasyxOpenapi_ShellGetSubmodelElements_Handler,
		},
		{
			MethodName: "ShellGetSubmodelElementByIdShort",
			Handler:    _BasyxOpenapi_ShellGetSubmodelElementByIdShort_Handler,
		},
		{
			MethodName: "ShellPutSubmodelElement",
			Handler:    _BasyxOpenapi_ShellPutSubmodelElement_Handler,
		},
		{
			MethodName: "ShellDeleteSubmodelElementByIdShort",
			Handler:    _BasyxOpenapi_ShellDeleteSubmodelElementByIdShort_Handler,
		},
		{
			MethodName: "ShellGetSubmodelElementValueByIdShort",
			Handler:    _BasyxOpenapi_ShellGetSubmodelElementValueByIdShort_Handler,
		},
		{
			MethodName: "ShellPutSubmodelElementValueByIdShort",
			Handler:    _BasyxOpenapi_ShellPutSubmodelElementValueByIdShort_Handler,
		},
		{
			MethodName: "ShellInvokeOperationByIdShort",
			Handler:    _BasyxOpenapi_ShellInvokeOperationByIdShort_Handler,
		},
		{
			MethodName: "ShellGetInvocationResultByIdShort",
			Handler:    _BasyxOpenapi_ShellGetInvocationResultByIdShort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gen/proto/basyx_openapi.proto",
}
