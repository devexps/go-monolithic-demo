package data

import (
	"context"
	"github.com/devexps/go-bootstrap"
	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/go-redis/redis/v8"

	conf "github.com/devexps/go-bootstrap/gen/api/go/conf/v1"
	"github.com/devexps/go-utils/entgo"

	"github.com/devexps/go-micro/middleware/authn/engine/jwt/v2"
	"github.com/devexps/go-micro/v2/middleware/authz/engine/noop"

	"github.com/devexps/go-micro/v2/log"

	authnEngine "github.com/devexps/go-micro/v2/middleware/authn/engine"
	authzEngine "github.com/devexps/go-micro/v2/middleware/authz/engine"

	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data/ent"
)

// Data .
type Data struct {
	log           *log.Helper
	db            *entgo.EntClient[*ent.Client]
	rdb           *redis.Client
	authenticator authnEngine.Authenticator
	authorizer    authzEngine.Authorizer
}

// NewData .
func NewData(
	logger log.Logger,
	db *entgo.EntClient[*ent.Client],
	rdb *redis.Client,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Authorizer,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/admin-service"))

	d := &Data{
		log:           l,
		db:            db,
		rdb:           rdb,
		authenticator: authenticator,
		authorizer:    authorizer,
	}
	return d, func() {
		l.Info("message", "closing the data resources")
		d.db.Close()
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewEntClient .
func NewEntClient(cfg *conf.Bootstrap, logger log.Logger) *entgo.EntClient[*ent.Client] {
	l := log.NewHelper(log.With(logger, "module", "ent/data/admin-service"))

	drv, err := entgo.CreateDriver(cfg.Data.Database.Driver, cfg.Data.Database.Source,
		int(cfg.Data.Database.MaxIdleConnections),
		int(cfg.Data.Database.MaxOpenConnections),
		cfg.Data.Database.ConnectionMaxLifetime.AsDuration(),
	)
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
		return nil
	}
	client := ent.NewClient(
		ent.Driver(drv),
		ent.Log(func(a ...any) {
			l.Debug(a...)
		}),
	)
	if cfg.Data.Database.Migrate {
		if err = client.Schema.Create(context.Background(), migrate.WithForeignKeys(true)); err != nil {
			l.Fatalf("failed creating schema resources: %v", err)
		}
	}
	return entgo.NewEntClient(client, drv)
}

// NewRedisClient creates redis client
func NewRedisClient(cfg *conf.Bootstrap, _ log.Logger) *redis.Client {
	return bootstrap.NewRedisClient(cfg.Data)
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
