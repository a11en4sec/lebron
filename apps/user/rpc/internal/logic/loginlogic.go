package logic

import (
	"context"
	"fmt"

	"github.com/a11en4sec/lebron/apps/user/rpc/internal/svc"
	"github.com/a11en4sec/lebron/apps/user/rpc/model"
	"github.com/a11en4sec/lebron/apps/user/rpc/user"
	"github.com/a11en4sec/lebron/pkg/tool"
	"github.com/a11en4sec/lebron/pkg/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {

	//verify user exists
	userDB, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		// 数据库没有错误,是没有找到记录
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "根据username查询用户信息失败,username:%s,err:%v", in.Username, err)
		}
	}

	// verify user password
	md5ByString, _ := tool.Md5ByString(in.Password)

	fmt.Println("[md5ByString]:", md5ByString)
	fmt.Println("userDB.Password:", userDB.Password)
	if !(md5ByString == userDB.Password) {
		return nil, errors.Wrap(xerr.NewErrMsg("账号或密码错误"), "密码错误")
	}

	//return sql
	var resp user.LoginResponse
	_ = copier.Copy(&resp, userDB)

	return &resp, nil

}
