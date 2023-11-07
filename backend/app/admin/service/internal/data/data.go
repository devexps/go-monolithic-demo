package data

import (
	conf "github.com/devexps/go-bootstrap/gen/api/go/conf/v1"
	"github.com/devexps/go-micro/middleware/authn/engine/jwt/v2"
	"github.com/devexps/go-micro/v2/middleware/authz/engine/noop"

	"github.com/devexps/go-micro/v2/log"

	authnEngine "github.com/devexps/go-micro/v2/middleware/authn/engine"
	authzEngine "github.com/devexps/go-micro/v2/middleware/authz/engine"
)

// Data .
type Data struct {
	log *log.Helper
}

// NewData .
func NewData(
	logger log.Logger,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/admin-service"))

	d := &Data{
		log: l,
	}
	return d, func() {
		l.Info("message", "closing the data resources")
	}, nil
}

// NewAuthenticator .
func NewAuthenticator(cfg *conf.Bootstrap) authnEngine.Authenticator {
	authenticator, _ := jwt.NewAuthenticator(
		jwt.WithKey([]byte(cfg.Server.Http.Middleware.Auth.Key)),
		jwt.WithSigningMethod(cfg.Server.Http.Middleware.Auth.Method),
	)
	return authenticator
}

// NewAuthorizer .
func NewAuthorizer() authzEngine.Authorizer {
	return noop.NewAuthorizer()
}
