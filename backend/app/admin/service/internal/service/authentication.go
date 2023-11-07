package service

import (
	"context"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data"

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
