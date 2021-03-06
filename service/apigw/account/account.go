package account

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"

	"micro/defined"
	"micro/pkg"
	"micro/response"
	"micro/service/account/proto"
)

type Account struct{}

var (
	accountCli proto.AccountService
)

func init() {
	service := micro.NewService(
		micro.Registry(pkg.RegistryConsul()),
	)
	service.Init()

	cli := service.Client()
	accountCli = proto.NewAccountService(
		defined.RPC_ACCOUNT_SERVICE_NAME, cli)
}

// 注册接口
func (account *Account) SinUpHandler(c *gin.Context) {
	var (
		req Register
		err error
	)

	// 参数校验
	err = pkg.ParseRequest(c, &req)
	if err != nil {
		return
	}

	// 登录流程调用rpc服务
	traceId, _ := c.Get(defined.TRACE_KEY)

	res, err := accountCli.SignIn(context.TODO(), &proto.SignInRequest{
		Username: req.Username,
		Password: req.Password,
		TranceId: traceId.(string),
	})

	if err != nil {
		c.Status(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	if res.Code == 10000 {
		c.JSON(http.StatusOK,
			response.NewErrorRes(res.Code, res.Message))
		return
	}

	c.JSON(http.StatusOK,
		response.NewSuccessRes(res.Code, res.Message, res.Data))
}

// 用户登陆接口
func (account *Account) SignInHandler(c *gin.Context) {
	var (
		req SignIn
		err error
	)

	err = pkg.ParseRequest(c, req)
	if err != nil {
		// TODO 记录系统运行异常日志
		return
	}

	traceId, _ := c.Get(defined.TRACE_KEY)
	res, err := accountCli.SignIn(context.TODO(), &proto.SignInRequest{
		Username: req.Username,
		Password: req.Password,
		Mobile: req.Mobile,
		VerifyCode: req.VerifyCode,
		TranceId: traceId.(string),
	})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if res.Code != 10000 {
		c.JSON(http.StatusOK,
			response.NewErrorRes(res.Code, res.Message))
	} else {
		c.JSON(http.StatusOK,
			response.NewSuccessRes(res.Code, res.Message, res.Data))
	}
}
