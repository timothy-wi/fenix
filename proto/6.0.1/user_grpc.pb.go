// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	Get(ctx context.Context, in *Authenticate, opts ...grpc.CallOption) (*User, error)
	New(ctx context.Context, in *CreateUser, opts ...grpc.CallOption) (*User, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

var usersGetStreamDesc = &grpc.StreamDesc{
	StreamName: "Get",
}

func (c *usersClient) Get(ctx context.Context, in *Authenticate, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/Users/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var usersNewStreamDesc = &grpc.StreamDesc{
	StreamName: "New",
}

func (c *usersClient) New(ctx context.Context, in *CreateUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/Users/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersService is the service API for Users service.
// Fields should be assigned to their respective handler implementations only before
// RegisterUsersService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type UsersService struct {
	Get func(context.Context, *Authenticate) (*User, error)
	New func(context.Context, *CreateUser) (*User, error)
}

func (s *UsersService) get(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Authenticate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/Users/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Get(ctx, req.(*Authenticate))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *UsersService) new(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/Users/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.New(ctx, req.(*CreateUser))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterUsersService registers a service implementation with a gRPC server.
func RegisterUsersService(s grpc.ServiceRegistrar, srv *UsersService) {
	srvCopy := *srv
	if srvCopy.Get == nil {
		srvCopy.Get = func(context.Context, *Authenticate) (*User, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
		}
	}
	if srvCopy.New == nil {
		srvCopy.New = func(context.Context, *CreateUser) (*User, error) {
			return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "Users",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Get",
				Handler:    srvCopy.get,
			},
			{
				MethodName: "New",
				Handler:    srvCopy.new,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "proto/6.0.1/user.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewUsersService creates a new UsersService containing the
// implemented methods of the Users service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewUsersService(s interface{}) *UsersService {
	ns := &UsersService{}
	if h, ok := s.(interface {
		Get(context.Context, *Authenticate) (*User, error)
	}); ok {
		ns.Get = h.Get
	}
	if h, ok := s.(interface {
		New(context.Context, *CreateUser) (*User, error)
	}); ok {
		ns.New = h.New
	}
	return ns
}

// UnstableUsersService is the service API for Users service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableUsersService interface {
	Get(context.Context, *Authenticate) (*User, error)
	New(context.Context, *CreateUser) (*User, error)
}
