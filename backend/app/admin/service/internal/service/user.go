package service

import (
	"github.com/devexps/go-micro/v2/log"

	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data"
	adminV1 "github.com/devexps/go-monolithic-demo/gen/api/go/admin/service/v1"
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
