package service

import (
	"github.com/devexps/go-micro/v2/log"

	adminV1 "github.com/devexps/go-monolithic-demo/gen/api/go/admin/service/v1"
)

type AuthenticationService struct {
	adminV1.AuthenticationServiceHTTPServer

	log *log.Helper
}

func NewAuthenticationService(logger log.Logger) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "authentication/service/admin-service"))
	return &AuthenticationService{
		log: l,
	}
}
