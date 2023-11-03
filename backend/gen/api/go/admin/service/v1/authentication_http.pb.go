// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             (unknown)
// source: admin/service/v1/authentication.proto

package v1

import (
	context "context"
	http "github.com/devexps/go-micro/v2/transport/http"
	binding "github.com/devexps/go-micro/v2/transport/http/binding"
	v1 "github.com/devexps/go-monolithic-demo/gen/api/go/user/service/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the micro package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAuthenticationServiceGetMe = "/admin.service.v1.AuthenticationService/GetMe"
const OperationAuthenticationServiceLogin = "/admin.service.v1.AuthenticationService/Login"
const OperationAuthenticationServiceLogout = "/admin.service.v1.AuthenticationService/Logout"
const OperationAuthenticationServiceRefreshToken = "/admin.service.v1.AuthenticationService/RefreshToken"

type AuthenticationServiceHTTPServer interface {
	// GetMe GetMe
	GetMe(context.Context, *GetMeRequest) (*v1.User, error)
	// Login Login
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Logout Logout
	Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error)
	// RefreshToken RefreshToken
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
}

func RegisterAuthenticationServiceHTTPServer(s *http.Server, srv AuthenticationServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/admin/v1/login", _AuthenticationService_Login0_HTTP_Handler(srv))
	r.POST("/admin/v1/logout", _AuthenticationService_Logout0_HTTP_Handler(srv))
	r.POST("/admin/v1/refresh_token", _AuthenticationService_RefreshToken0_HTTP_Handler(srv))
	r.GET("/admin/v1/me", _AuthenticationService_GetMe0_HTTP_Handler(srv))
}

func _AuthenticationService_Login0_HTTP_Handler(srv AuthenticationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthenticationServiceLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResponse)
		return ctx.Result(200, reply)
	}
}

func _AuthenticationService_Logout0_HTTP_Handler(srv AuthenticationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthenticationServiceLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _AuthenticationService_RefreshToken0_HTTP_Handler(srv AuthenticationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RefreshTokenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthenticationServiceRefreshToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RefreshToken(ctx, req.(*RefreshTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RefreshTokenResponse)
		return ctx.Result(200, reply)
	}
}

func _AuthenticationService_GetMe0_HTTP_Handler(srv AuthenticationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthenticationServiceGetMe)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMe(ctx, req.(*GetMeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.User)
		return ctx.Result(200, reply)
	}
}

type AuthenticationServiceHTTPClient interface {
	GetMe(ctx context.Context, req *GetMeRequest, opts ...http.CallOption) (rsp *v1.User, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginResponse, err error)
	Logout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	RefreshToken(ctx context.Context, req *RefreshTokenRequest, opts ...http.CallOption) (rsp *RefreshTokenResponse, err error)
}

type AuthenticationServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthenticationServiceHTTPClient(client *http.Client) AuthenticationServiceHTTPClient {
	return &AuthenticationServiceHTTPClientImpl{client}
}

func (c *AuthenticationServiceHTTPClientImpl) GetMe(ctx context.Context, in *GetMeRequest, opts ...http.CallOption) (*v1.User, error) {
	var out v1.User
	pattern := "/admin/v1/me"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthenticationServiceGetMe))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthenticationServiceHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginResponse, error) {
	var out LoginResponse
	pattern := "/admin/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthenticationServiceLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthenticationServiceHTTPClientImpl) Logout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthenticationServiceLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthenticationServiceHTTPClientImpl) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...http.CallOption) (*RefreshTokenResponse, error) {
	var out RefreshTokenResponse
	pattern := "/admin/v1/refresh_token"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthenticationServiceRefreshToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}