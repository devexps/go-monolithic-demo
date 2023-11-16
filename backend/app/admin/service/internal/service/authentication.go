package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data"
	"github.com/devexps/go-monolithic-demo/pkg/middleware/auth"

	adminV1 "github.com/devexps/go-monolithic-demo/gen/api/go/admin/service/v1"
	userV1 "github.com/devexps/go-monolithic-demo/gen/api/go/user/service/v1"
)

// AuthenticationService .
type AuthenticationService struct {
	adminV1.AuthenticationServiceHTTPServer

	log  *log.Helper
	uc   *data.UserRepo
	utuc *data.UserTokenRepo
}

// NewAuthenticationService .
func NewAuthenticationService(logger log.Logger, uc *data.UserRepo, utuc *data.UserTokenRepo) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "authentication/service/admin-service"))
	return &AuthenticationService{
		log:  l,
		uc:   uc,
		utuc: utuc,
	}
}

// Login .
func (s *AuthenticationService) Login(ctx context.Context, req *adminV1.LoginRequest) (*adminV1.LoginResponse, error) {
	var user *userV1.User
	var err error
	if user, err = s.uc.VerifyPassword(ctx, &userV1.VerifyPasswordRequest{
		UserName: req.GetUsername(),
		Password: req.GetPassword(),
	}); err != nil {
		return nil, err
	}
	tokenRsp, err := s.utuc.GenerateToken(ctx, user)
	if err != nil {
		return nil, err
	}
	return &adminV1.LoginResponse{
		GrandType:    "bearer",
		AccessToken:  tokenRsp.GetAccessToken(),
		RefreshToken: tokenRsp.GetRefreshToken(),
	}, nil
}

// Logout .
func (s *AuthenticationService) Logout(ctx context.Context, req *adminV1.LogoutRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	req.Id = authInfo.UserId

	_, err = s.utuc.RemoveToken(ctx, &userV1.User{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// RefreshToken .
func (s *AuthenticationService) RefreshToken(ctx context.Context, req *adminV1.RefreshTokenRequest) (*adminV1.RefreshTokenResponse, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	rsp, err := s.utuc.GetRefreshToken(ctx, &userV1.User{
		Id: authInfo.UserId,
	})
	if err != nil {
		return nil, err
	}
	if rsp.GetRefreshToken() != req.GetRefreshToken() {
		return nil, adminV1.ErrorInvalidToken("invalid refresh token")
	}
	rsp2, err := s.utuc.GenerateAccessToken(ctx, &userV1.User{
		Id: authInfo.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &adminV1.RefreshTokenResponse{
		GrandType:    "bearer",
		RefreshToken: rsp.GetRefreshToken(),
		AccessToken:  rsp2.GetAccessToken(),
	}, nil
}

// GetMe .
func (s *AuthenticationService) GetMe(ctx context.Context, req *adminV1.GetMeRequest) (*userV1.User, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	req.Id = authInfo.UserId

	userRsp, err := s.uc.GetUser(ctx, &userV1.GetUserRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return userRsp, nil
}
