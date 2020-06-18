package modules

import (
	"context"
	"micro/dao"
	"micro/defined"
	"micro/service/account/proto"
	"micro/service/account/validate"
	"micro/util"
	"time"
)

type Account struct{}

func(u *Account)SignIn(ctx context.Context,
	req *proto.SignInRequest, res *proto.SignInResponse) error {
	// 参数校验
	if req.Mobile == "" && req.Username == "" {
		res.Code = 10005
		res.Message = "请填写用户名或手机号"
		return nil
	}
	tx := context.WithValue(ctx, defined.TRACE_KEY, req.TranceId)

	user, err := (&dao.User{}).Find(tx, 1)
	if err != nil {
		return err
	}

	if user.Password == req.Password {
		res.Code = 10000
		res.Message = "登陆成功"
		res.Data = &proto.SignInResponse_SuccessData{
			Token: util.GenerateToken(user.Mobile),
		}
	} else {
		res.Code = 10006
		res.Message = "用户名或密码错误"
		return nil
	}

	return nil
}

func (u *Account) SignUp(
	ctx context.Context, req *proto.SignUpRequest, res *proto.SignUpResponse) error {
	// 参数校验
	err := validate.ValidatorRegister(req)
	if err != nil {
		res.Code = 10005
		res.Message = err.Error()
		return nil
	}

	tx := context.WithValue(ctx, defined.TRACE_KEY, req.TranceId)
	user := &dao.User{
		Username: req.Username,
		Password: req.Password,
		Mobile: req.Mobile,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = (&dao.User{}).Create(tx, user)
	if err != nil {
		res.Code = 10006
		res.Message = err.Error()
	} else {
		res.Code = 10000
		res.Message = "添加用户成功"
	}
	return nil
}
