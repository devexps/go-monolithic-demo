package data

import (
	"context"
	"github.com/devexps/go-utils/crypto"

	"github.com/devexps/go-micro/v2/log"

	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data/ent"
	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data/ent/user"
	v1 "github.com/devexps/go-monolithic-demo/gen/api/go/user/service/v1"
)

// UserRepo .
type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/repo/admin-service"))
	return &UserRepo{
		data: data,
		log:  l,
	}
}

// VerifyPassword .
func (r *UserRepo) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.User, error) {
	resp, err := r.data.db.Client().User.
		Query().
		Select(user.FieldID, user.FieldUserName, user.FieldPassword).
		Where(user.UserNameEQ(req.GetUserName())).
		Only(ctx)
	if err != nil {
		r.log.Errorf("VerifyPassword query user data failed, err=%v", err)
		return nil, v1.ErrorUserNotExist("user not exist")
	}
	bMatched := crypto.CheckPasswordHash(req.GetPassword(), resp.Password)
	if !bMatched {
		return nil, v1.ErrorInvalidPassword("invalid password")
	}
	u := r.convertEntToProto(resp)
	return u, err
}

// GetUser .
func (r *UserRepo) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	resp, err := r.data.db.Client().User.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, v1.ErrorUserNotExist("user not found")
		}
		r.log.Errorf("GetUser query one data failed, err=%v", err)
		return nil, v1.ErrorQueryDataFailed("query data failed")
	}
	u := r.convertEntToProto(resp)
	return u, err
}

func (r *UserRepo) convertEntToProto(in *ent.User) *v1.User {
	if in == nil {
		return nil
	}
	return &v1.User{
		Id:       in.ID,
		UserName: in.UserName,
		NickName: in.NickName,
		RealName: in.RealName,
		Email:    in.Email,
		Phone:    in.Phone,
	}
}
