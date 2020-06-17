package modules

import (
	"context"
	"micro/dao"
	"micro/db"
	"micro/defind"
	"micro/service/account/proto"
	"micro/util"
)

type Account struct{}

func(u *Account)SinIn(ctx context.Context,
	req *proto.SinInRequest, res *proto.SinInResponse) error {
	// 参数校验
	if req.Mobile == "" && req.Username == "" {
		res.Code = 10005
		res.Message = "请填写用户名或手机号"
		return nil
	}
	tx := context.WithValue(ctx, defind.TRACE_KEY, req.TranceId)
	gorm, err := db.GetMasterDBConn(tx)
	if err != nil {
		return err
	}


	user, err := (&dao.User{}).Find(gorm, 1)
	if err != nil {
		return err
	}

	if user.Password == req.Password {
		res.Code = 10000
		res.Message = "登陆成功"
		res.Data = &proto.SinInResponse_SuccessData{
			Token: util.GenerateToken(user.Mobile),
		}
	} else {
		res.Code = 10006
		res.Message = "用户名或密码错误"
		return nil
	}

	return nil
}
