// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: data.proto

package protodata

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
	PersonRepository_CreatePerson_FullMethodName    = "/protodata.PersonRepository/CreatePerson"
	PersonRepository_UpdatePerson_FullMethodName    = "/protodata.PersonRepository/UpdatePerson"
	PersonRepository_GetPersonByGuid_FullMethodName = "/protodata.PersonRepository/GetPersonByGuid"
	PersonRepository_FindAllPeople_FullMethodName   = "/protodata.PersonRepository/FindAllPeople"
)

// PersonRepositoryClient is the client API for PersonRepository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonRepositoryClient interface {
	CreatePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error)
	UpdatePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error)
	GetPersonByGuid(ctx context.Context, in *PersonByGuidRequest, opts ...grpc.CallOption) (*Person, error)
	FindAllPeople(ctx context.Context, in *Filters, opts ...grpc.CallOption) (*PersonList, error)
}

type personRepositoryClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonRepositoryClient(cc grpc.ClientConnInterface) PersonRepositoryClient {
	return &personRepositoryClient{cc}
}

func (c *personRepositoryClient) CreatePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Person)
	err := c.cc.Invoke(ctx, PersonRepository_CreatePerson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personRepositoryClient) UpdatePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Person)
	err := c.cc.Invoke(ctx, PersonRepository_UpdatePerson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personRepositoryClient) GetPersonByGuid(ctx context.Context, in *PersonByGuidRequest, opts ...grpc.CallOption) (*Person, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Person)
	err := c.cc.Invoke(ctx, PersonRepository_GetPersonByGuid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personRepositoryClient) FindAllPeople(ctx context.Context, in *Filters, opts ...grpc.CallOption) (*PersonList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PersonList)
	err := c.cc.Invoke(ctx, PersonRepository_FindAllPeople_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonRepositoryServer is the server API for PersonRepository service.
// All implementations must embed UnimplementedPersonRepositoryServer
// for forward compatibility.
type PersonRepositoryServer interface {
	CreatePerson(context.Context, *Person) (*Person, error)
	UpdatePerson(context.Context, *Person) (*Person, error)
	GetPersonByGuid(context.Context, *PersonByGuidRequest) (*Person, error)
	FindAllPeople(context.Context, *Filters) (*PersonList, error)
	mustEmbedUnimplementedPersonRepositoryServer()
}

// UnimplementedPersonRepositoryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPersonRepositoryServer struct{}

func (UnimplementedPersonRepositoryServer) CreatePerson(context.Context, *Person) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (UnimplementedPersonRepositoryServer) UpdatePerson(context.Context, *Person) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (UnimplementedPersonRepositoryServer) GetPersonByGuid(context.Context, *PersonByGuidRequest) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonByGuid not implemented")
}
func (UnimplementedPersonRepositoryServer) FindAllPeople(context.Context, *Filters) (*PersonList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllPeople not implemented")
}
func (UnimplementedPersonRepositoryServer) mustEmbedUnimplementedPersonRepositoryServer() {}
func (UnimplementedPersonRepositoryServer) testEmbeddedByValue()                          {}

// UnsafePersonRepositoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonRepositoryServer will
// result in compilation errors.
type UnsafePersonRepositoryServer interface {
	mustEmbedUnimplementedPersonRepositoryServer()
}

func RegisterPersonRepositoryServer(s grpc.ServiceRegistrar, srv PersonRepositoryServer) {
	// If the following call pancis, it indicates UnimplementedPersonRepositoryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PersonRepository_ServiceDesc, srv)
}

func _PersonRepository_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonRepositoryServer).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonRepository_CreatePerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonRepositoryServer).CreatePerson(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonRepository_UpdatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonRepositoryServer).UpdatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonRepository_UpdatePerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonRepositoryServer).UpdatePerson(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonRepository_GetPersonByGuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonByGuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonRepositoryServer).GetPersonByGuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonRepository_GetPersonByGuid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonRepositoryServer).GetPersonByGuid(ctx, req.(*PersonByGuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonRepository_FindAllPeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonRepositoryServer).FindAllPeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonRepository_FindAllPeople_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonRepositoryServer).FindAllPeople(ctx, req.(*Filters))
	}
	return interceptor(ctx, in, info, handler)
}

// PersonRepository_ServiceDesc is the grpc.ServiceDesc for PersonRepository service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonRepository_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protodata.PersonRepository",
	HandlerType: (*PersonRepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePerson",
			Handler:    _PersonRepository_CreatePerson_Handler,
		},
		{
			MethodName: "UpdatePerson",
			Handler:    _PersonRepository_UpdatePerson_Handler,
		},
		{
			MethodName: "GetPersonByGuid",
			Handler:    _PersonRepository_GetPersonByGuid_Handler,
		},
		{
			MethodName: "FindAllPeople",
			Handler:    _PersonRepository_FindAllPeople_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}
