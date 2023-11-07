package data

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"

	"github.com/devexps/go-micro/v2/log"

	"github.com/devexps/go-monolithic-demo/gen/api/go/common/authn"
	userV1 "github.com/devexps/go-monolithic-demo/gen/api/go/user/service/v1"
	v1 "github.com/devexps/go-monolithic-demo/gen/api/go/user_token/service/v1"
)

const (
	userAccessTokenKeyPrefix  = "a_uat_"
	userRefreshTokenKeyPrefix = "a_urt_"
)

// UserTokenRepo .
type UserTokenRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserTokenRepo .
func NewUserTokenRepo(data *Data, logger log.Logger) *UserTokenRepo {
	l := log.NewHelper(log.With(logger, "module", "user_token/repo/admin-service"))
	return &UserTokenRepo{
		data: data,
		log:  l,
	}
}

// GenerateToken .
func (r *UserTokenRepo) GenerateToken(ctx context.Context, user *userV1.User) (*v1.TokenResponse, error) {
	accessToken := r.createAccessJwtToken(user.GetId(), user.GetUserName())
	if accessToken == "" {
		return nil, v1.ErrorCreateAccessTokenFailed("create access token failed")
	}
	if err := r.setAccessTokenToRedis(ctx, user.GetId(), accessToken, 0); err != nil {
		r.log.Errorf("GenerateToken store to redis failed, err=%v", err)
		return nil, v1.ErrorStoreRedisFailed("store to redis failed")
	}
	refreshToken := r.createRefreshToken()
	if refreshToken == "" {
		return nil, v1.ErrorCreateRefreshTokenFailed("create refresh token failed")
	}
	if err := r.setRefreshTokenToRedis(ctx, user.GetId(), refreshToken, 0); err != nil {
		r.log.Errorf("GenerateToken store to redis failed, err=%v", err)
		return nil, v1.ErrorStoreRedisFailed("store to redis failed")
	}
	return &v1.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (r *UserTokenRepo) createAccessJwtToken(userId, userName string) string {
	principal := &authn.Claims{
		Subject:  userName,
		UserId:   userId,
		UserName: userName,
	}
	signedToken, err := r.data.authenticator.CreateIdentity(principal)
	if err != nil {
		return ""
	}
	return signedToken
}

func (r *UserTokenRepo) createRefreshToken() string {
	strUUID, _ := uuid.NewV4()
	return strUUID.String()
}

func (r *UserTokenRepo) setAccessTokenToRedis(ctx context.Context, userId, token string, expires int32) error {
	key := fmt.Sprintf("%s%s", userAccessTokenKeyPrefix, userId)
	return r.data.rdb.Set(ctx, key, token, time.Duration(expires)).Err()
}

func (r *UserTokenRepo) setRefreshTokenToRedis(ctx context.Context, userId, token string, expires int32) error {
	key := fmt.Sprintf("%s%s", userRefreshTokenKeyPrefix, userId)
	return r.data.rdb.Set(ctx, key, token, time.Duration(expires)).Err()
}
