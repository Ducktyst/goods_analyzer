// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// PriceAnalyzeClient is the client API for PriceAnalyze service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PriceAnalyzeClient interface {
	// список товаров с фильтром по категорией
	GetProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductListResponse, error)
	// список категорий
	GetCategoryList(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*CategoryListResponse, error)
	GetProductsFilter(ctx context.Context, in *GetProductsFilterRequest, opts ...grpc.CallOption) (*GetProductsFilterResponse, error)
	AddProductsFilter(ctx context.Context, in *AddProductsFilterRequest, opts ...grpc.CallOption) (*AddProductsFilterResponse, error)
}

type priceAnalyzeClient struct {
	cc grpc.ClientConnInterface
}

func NewPriceAnalyzeClient(cc grpc.ClientConnInterface) PriceAnalyzeClient {
	return &priceAnalyzeClient{cc}
}

func (c *priceAnalyzeClient) GetProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductListResponse, error) {
	out := new(ProductListResponse)
	err := c.cc.Invoke(ctx, "/api.PriceAnalyze/GetProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceAnalyzeClient) GetCategoryList(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*CategoryListResponse, error) {
	out := new(CategoryListResponse)
	err := c.cc.Invoke(ctx, "/api.PriceAnalyze/GetCategoryList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceAnalyzeClient) GetProductsFilter(ctx context.Context, in *GetProductsFilterRequest, opts ...grpc.CallOption) (*GetProductsFilterResponse, error) {
	out := new(GetProductsFilterResponse)
	err := c.cc.Invoke(ctx, "/api.PriceAnalyze/GetProductsFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceAnalyzeClient) AddProductsFilter(ctx context.Context, in *AddProductsFilterRequest, opts ...grpc.CallOption) (*AddProductsFilterResponse, error) {
	out := new(AddProductsFilterResponse)
	err := c.cc.Invoke(ctx, "/api.PriceAnalyze/AddProductsFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceAnalyzeServer is the server API for PriceAnalyze service.
// All implementations must embed UnimplementedPriceAnalyzeServer
// for forward compatibility
type PriceAnalyzeServer interface {
	// список товаров с фильтром по категорией
	GetProductList(context.Context, *ProductListRequest) (*ProductListResponse, error)
	// список категорий
	GetCategoryList(context.Context, *CategoryListRequest) (*CategoryListResponse, error)
	GetProductsFilter(context.Context, *GetProductsFilterRequest) (*GetProductsFilterResponse, error)
	AddProductsFilter(context.Context, *AddProductsFilterRequest) (*AddProductsFilterResponse, error)
	mustEmbedUnimplementedPriceAnalyzeServer()
}

// UnimplementedPriceAnalyzeServer must be embedded to have forward compatible implementations.
type UnimplementedPriceAnalyzeServer struct {
}

func (UnimplementedPriceAnalyzeServer) GetProductList(context.Context, *ProductListRequest) (*ProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductList not implemented")
}
func (UnimplementedPriceAnalyzeServer) GetCategoryList(context.Context, *CategoryListRequest) (*CategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryList not implemented")
}
func (UnimplementedPriceAnalyzeServer) GetProductsFilter(context.Context, *GetProductsFilterRequest) (*GetProductsFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsFilter not implemented")
}
func (UnimplementedPriceAnalyzeServer) AddProductsFilter(context.Context, *AddProductsFilterRequest) (*AddProductsFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProductsFilter not implemented")
}
func (UnimplementedPriceAnalyzeServer) mustEmbedUnimplementedPriceAnalyzeServer() {}

// UnsafePriceAnalyzeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PriceAnalyzeServer will
// result in compilation errors.
type UnsafePriceAnalyzeServer interface {
	mustEmbedUnimplementedPriceAnalyzeServer()
}

func RegisterPriceAnalyzeServer(s grpc.ServiceRegistrar, srv PriceAnalyzeServer) {
	s.RegisterService(&PriceAnalyze_ServiceDesc, srv)
}

func _PriceAnalyze_GetProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceAnalyzeServer).GetProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PriceAnalyze/GetProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceAnalyzeServer).GetProductList(ctx, req.(*ProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceAnalyze_GetCategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceAnalyzeServer).GetCategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PriceAnalyze/GetCategoryList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceAnalyzeServer).GetCategoryList(ctx, req.(*CategoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceAnalyze_GetProductsFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceAnalyzeServer).GetProductsFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PriceAnalyze/GetProductsFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceAnalyzeServer).GetProductsFilter(ctx, req.(*GetProductsFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceAnalyze_AddProductsFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductsFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceAnalyzeServer).AddProductsFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PriceAnalyze/AddProductsFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceAnalyzeServer).AddProductsFilter(ctx, req.(*AddProductsFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PriceAnalyze_ServiceDesc is the grpc.ServiceDesc for PriceAnalyze service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PriceAnalyze_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.PriceAnalyze",
	HandlerType: (*PriceAnalyzeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductList",
			Handler:    _PriceAnalyze_GetProductList_Handler,
		},
		{
			MethodName: "GetCategoryList",
			Handler:    _PriceAnalyze_GetCategoryList_Handler,
		},
		{
			MethodName: "GetProductsFilter",
			Handler:    _PriceAnalyze_GetProductsFilter_Handler,
		},
		{
			MethodName: "AddProductsFilter",
			Handler:    _PriceAnalyze_AddProductsFilter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/price_analyze.proto",
}
