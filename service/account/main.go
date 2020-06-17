package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"micro/bootstrap"
	"micro/defind"
	"micro/pkg"
	"time"

	"micro/service/account/modules"
	"micro/service/account/proto"
)

func main() {
	service := micro.NewService(
		micro.Registry(pkg.RegistryConsul()),// 注册consul
		micro.Name(defind.RPC_ACCOUNT_SERVICE_NAME),
		micro.RegisterTTL(time.Second * 10),
		micro.RegisterInterval(time.Second * 5),
	)
	service.Init()

	err := proto.RegisterAccountServiceHandler(service.Server(), new(modules.Account))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = bootstrap.InitModule("../../conf/dev")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
