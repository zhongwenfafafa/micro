package modules

import (
	"context"
	"fmt"
	"micro/bootstrap"
	"micro/pkg"
	"micro/service/account/proto"
	"testing"
)

func TestAccountSignIn(t *testing.T) {
	SetUp()
	account := &Account{}
	in := &proto.SignInRequest{
		Username: "zhongwen",
		Password: "123456",
		TranceId: pkg.GenerateUUID(),
	}

	out := &proto.SignInResponse{}
	err := account.SignIn(context.TODO(), in, out)
	if err != nil {
		t.Errorf("sign in failed, err:%s", err)
	}

	if out.Code != 10000 {
		t.Errorf("sign in failed, err code:%d, err " +
			"message:%s", out.Code, out.Message)
	} else {
		t.Log(out.Code, out.Message)
	}
}

func TestAccountSignUp(t *testing.T) {
	SetUp()
	account := &Account{}

	in := &proto.SignUpRequest{
		Username: "zhongwen1",
		Password: "123456",
		Mobile: "17756309908",
		TranceId: pkg.GenerateUUID(),
	}

	out := &proto.SignUpResponse{}

	err := account.SignUp(context.TODO(), in, out)
	if err != nil {
		t.Errorf("sign in failed, err:%s", err)
	}

	if out.Code != 10000 {
		t.Errorf("sign in failed, err code:%d, err " +
			"message:%s", out.Code, out.Message)
	} else {
		t.Log(out.Code, out.Message)
	}
}

func SetUp() {
	err := bootstrap.InitModule("../../../conf/dev")
	if err != nil {
		fmt.Printf("start module failed, err:%s", err)
		return
	}
}