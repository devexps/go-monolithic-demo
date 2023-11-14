package service

import (
	"context"

	"github.com/devexps/go-micro/v2/log"

	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data"
	adminV1 "github.com/devexps/go-monolithic-demo/gen/api/go/admin/service/v1"
	"github.com/devexps/go-monolithic-demo/gen/api/go/common/pagination"
	v1 "github.com/devexps/go-monolithic-demo/gen/api/go/user/service/v1"
)

type UserService struct {
	adminV1.UserServiceHTTPServer

	log *log.Helper
	uc  *data.UserRepo
}

func NewUserService(logger log.Logger, uc *data.UserRepo) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/admin-service"))
	return &UserService{
		log: l,
		uc:  uc,
	}
}

func (s *UserService) ListUser(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	return s.uc.ListUser(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	return s.uc.GetUser(ctx, req)
}
