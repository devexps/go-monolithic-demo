package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	entgoUpdate "github.com/devexps/go-utils/entgo/update"
	"github.com/devexps/go-utils/fieldmask"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"

	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-utils/crypto"
	entgoQuery "github.com/devexps/go-utils/entgo/query"

	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data/ent"
	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data/ent/user"
	"github.com/devexps/go-monolithic-demo/gen/api/go/common/pagination"
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
		Select(user.FieldID, user.FieldUsername, user.FieldPassword).
		Where(user.UsernameEQ(req.GetUserName())).
		Only(ctx)
	if err != nil {
		r.log.Errorf("VerifyPassword query user data failed, err=%v", err)
		return nil, v1.ErrorUserNotExist("user not exist")
	}
	bMatched := crypto.CheckPasswordHash(req.GetPassword(), *resp.Password)
	if !bMatched {
		//hintPass, _ := crypto.HashPassword(req.GetPassword())
		//r.log.Warn(hintPass)
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

// ListUser .
func (r *UserRepo) ListUser(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	builder := r.data.db.Client().User.Query()

	err, whereSelectors, querySelectors := entgoQuery.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), user.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("ListUser build query failed: %v", err)
		return nil, v1.ErrorQueryDataFailed("build query failed")
	}
	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}
	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("ListUser query list failed: %v", err)
		return nil, v1.ErrorQueryDataFailed("query list failed")
	}
	items := make([]*v1.User, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}
	count, err := r.count(ctx, whereSelectors)
	if err != nil {
		r.log.Errorf("ListUser count rows failed: %v", err)
		return nil, v1.ErrorQueryDataFailed("count data failed")
	}
	return &v1.ListUserResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

// CreateUser .
func (r *UserRepo) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*emptypb.Empty, error) {
	builder := r.data.db.Client().User.Create().
		SetNillableUsername(req.User.UserName).
		SetNillableNickName(req.User.NickName).
		SetNillableRealName(req.User.RealName).
		SetNillableEmail(req.User.Email).
		SetNillablePhone(req.User.Phone).
		SetNillableStatus((*user.Status)(req.User.Status)).
		SetCreateBy(req.GetOperatorId()).
		SetCreateTime(time.Now())

	if req.User.Authority != nil {
		builder.SetAuthority((user.Authority)(req.User.Authority.String()))
	}
	if (len(req.Password)) > 0 {
		cryptoPassword, err := crypto.HashPassword(req.GetPassword())
		if err == nil {
			builder.SetPassword(cryptoPassword)
		}
	}
	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("CreateUser insert one data failed: %v", err)
		return nil, v1.ErrorInsertDataFailed("insert one data failed")
	}
	return &emptypb.Empty{}, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*emptypb.Empty, error) {
	if req.GetUpdateMask() == nil {
		return nil, v1.ErrorInvalidRequest("invalid field mask")
	}
	req.UpdateMask.Normalize()

	if !req.UpdateMask.IsValid(req.User) {
		return nil, v1.ErrorInvalidFieldMask("invalid field mask")
	}
	fieldmask.Filter(req.GetUser(), req.UpdateMask.GetPaths())

	builder := r.data.db.Client().User.UpdateOneID(req.User.Id).
		SetNillableNickName(req.User.NickName).
		SetNillableRealName(req.User.RealName).
		SetNillableEmail(req.User.Email).
		SetNillablePhone(req.User.Phone).
		SetNillableStatus((*user.Status)(req.User.Status)).
		SetUpdateTime(time.Now())

	if req.User.Authority != nil {
		builder.SetAuthority((user.Authority)(req.User.Authority.String()))
	}
	if len(req.Password) > 0 {
		cryptoPassword, err := crypto.HashPassword(req.GetPassword())
		if err == nil {
			builder.SetPassword(cryptoPassword)
		}
	}
	nilPaths := fieldmask.NilValuePaths(req.User, req.GetUpdateMask().GetPaths())
	nilUpdater, _ := entgoUpdate.BuildSetNullUpdater(nilPaths)
	if nilUpdater != nil {
		builder.Modify(nilUpdater)
	}
	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("UpdateUser update one data failed: %v", err)
		return nil, v1.ErrorUpdateDataFailed("update one data failed")
	}
	return nil, nil
}

func (r *UserRepo) convertEntToProto(in *ent.User) *v1.User {
	if in == nil {
		return nil
	}
	return &v1.User{
		Id:       in.ID,
		UserName: in.Username,
		NickName: in.NickName,
		RealName: in.RealName,
		Email:    in.Email,
		Phone:    in.Phone,
	}
}

func (r *UserRepo) count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().User.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}
	return builder.Count(ctx)
}
