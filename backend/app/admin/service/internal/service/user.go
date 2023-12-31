package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-utils/trans"

	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data"
	adminV1 "github.com/devexps/go-monolithic-demo/gen/api/go/admin/service/v1"
	"github.com/devexps/go-monolithic-demo/gen/api/go/common/pagination"
	v1 "github.com/devexps/go-monolithic-demo/gen/api/go/user/service/v1"
	"github.com/devexps/go-monolithic-demo/pkg/middleware/auth"
)

// UserService .
type UserService struct {
	adminV1.UserServiceHTTPServer

	log *log.Helper
	uc  *data.UserRepo
}

// NewUserService .
func NewUserService(logger log.Logger, uc *data.UserRepo) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/admin-service"))
	return &UserService{
		log: l,
		uc:  uc,
	}
}

// ListUser .
func (s *UserService) ListUser(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	return s.uc.ListUser(ctx, req)
}

// GetUser .
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	return s.uc.GetUser(ctx, req)
}

// CreateUser .
func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	if req.User == nil {
		return nil, adminV1.ErrorBadRequest("user is nil")
	}
	req.OperatorId = authInfo.UserId
	req.User.CreatorId = trans.Uint32(authInfo.UserId)

	if req.User.Authority == nil {
		req.User.Authority = v1.UserAuthority_CUSTOMER_USER.Enum()
	}
	if _, err = s.uc.CreateUser(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// UpdateUser .
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	if req.User == nil {
		return nil, adminV1.ErrorBadRequest("user is nil")
	}
	req.OperatorId = authInfo.UserId

	if _, err := s.uc.UpdateUser(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// DeleteUser .
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	req.OperatorId = authInfo.UserId

	if _, err := s.uc.DeleteUser(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
