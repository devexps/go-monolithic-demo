// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.0.6
// - protoc             (unknown)
// source: admin/service/v1/user.proto

package v1

import (
	context "context"
	http "github.com/devexps/go-micro/v2/transport/http"
	binding "github.com/devexps/go-micro/v2/transport/http/binding"
	pagination "github.com/devexps/go-monolithic-demo/gen/api/go/common/pagination"
	v1 "github.com/devexps/go-monolithic-demo/gen/api/go/user/service/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the micro package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserServiceCreateUser = "/admin.service.v1.UserService/CreateUser"
const OperationUserServiceDeleteUser = "/admin.service.v1.UserService/DeleteUser"
const OperationUserServiceGetUser = "/admin.service.v1.UserService/GetUser"
const OperationUserServiceListUser = "/admin.service.v1.UserService/ListUser"
const OperationUserServiceUpdateUser = "/admin.service.v1.UserService/UpdateUser"

type UserServiceHTTPServer interface {
	// CreateUser CreateUser
	CreateUser(context.Context, *v1.CreateUserRequest) (*emptypb.Empty, error)
	// DeleteUser DeleteUser
	DeleteUser(context.Context, *v1.DeleteUserRequest) (*emptypb.Empty, error)
	// GetUser GetUser
	GetUser(context.Context, *v1.GetUserRequest) (*v1.User, error)
	// ListUser ListUser
	ListUser(context.Context, *pagination.PagingRequest) (*v1.ListUserResponse, error)
	// UpdateUser UpdateUser
	UpdateUser(context.Context, *v1.UpdateUserRequest) (*emptypb.Empty, error)
}

func RegisterUserServiceHTTPServer(s *http.Server, srv UserServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/users", _UserService_ListUser0_HTTP_Handler(srv))
	r.GET("/admin/v1/users/{id}", _UserService_GetUser0_HTTP_Handler(srv))
	r.POST("/admin/v1/users", _UserService_CreateUser0_HTTP_Handler(srv))
	r.PUT("/admin/v1/users/{user.id}", _UserService_UpdateUser0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/users/{id}", _UserService_DeleteUser0_HTTP_Handler(srv))
}

func _UserService_ListUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in pagination.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*pagination.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.ListUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*v1.GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.User)
		return ctx.Result(200, reply)
	}
}

func _UserService_CreateUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.CreateUserRequest
		if err := ctx.Bind(&in.User); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*v1.CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _UserService_UpdateUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UpdateUserRequest
		if err := ctx.Bind(&in.User); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*v1.UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _UserService_DeleteUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.DeleteUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*v1.DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type UserServiceHTTPClient interface {
	CreateUser(ctx context.Context, req *v1.CreateUserRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteUser(ctx context.Context, req *v1.DeleteUserRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetUser(ctx context.Context, req *v1.GetUserRequest, opts ...http.CallOption) (rsp *v1.User, err error)
	ListUser(ctx context.Context, req *pagination.PagingRequest, opts ...http.CallOption) (rsp *v1.ListUserResponse, err error)
	UpdateUser(ctx context.Context, req *v1.UpdateUserRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type UserServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserServiceHTTPClient(client *http.Client) UserServiceHTTPClient {
	return &UserServiceHTTPClientImpl{client}
}

func (c *UserServiceHTTPClientImpl) CreateUser(ctx context.Context, in *v1.CreateUserRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.User, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) GetUser(ctx context.Context, in *v1.GetUserRequest, opts ...http.CallOption) (*v1.User, error) {
	var out v1.User
	pattern := "/admin/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) ListUser(ctx context.Context, in *pagination.PagingRequest, opts ...http.CallOption) (*v1.ListUserResponse, error) {
	var out v1.ListUserResponse
	pattern := "/admin/v1/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/users/{user.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.User, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
