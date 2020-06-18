package db

import (
	"context"
	"flag"
	"fmt"
	"log"
	"micro/dao"
	"micro/defined"
	"micro/pkg"
	"os"
	"testing"
	"time"
)

func TestGetMasterDBConn(t *testing.T) {
	SetUp("../conf/dev")
	err := InitDBPool()
	if err != nil {
		t.Error(err)
	}

	ctx := context.WithValue(context.TODO(), "traceId", pkg.GenerateUUID())
	conn, err := GetMasterDBConn(ctx)
	if err != nil {
		t.Errorf("get connection err, %s", err)
	}

	account := &dao.User{
		Username:  "zhongwen",
		Password:  "123456",
		Mobile:    "17756309909",
		IsDeleted: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	conn.Table("account").Find(&account).Where("id=?", 1)
	fmt.Println(account.Id)
}

func SetUp(confPath string) {
	conf := flag.String("config", confPath, "input config file like ./conf/dev/")
	flag.Parse()

	if *conf == "" {
		flag.Usage()
		os.Exit(1)
	}

	log.Println("------------------------------------------------------------------------")
	log.Printf("[INFO]  config=%s\n", *conf)
	log.Printf("[INFO] %s\n", " start loading resources.")

	// 解析配置文件目录
	if err := pkg.ParseConfPath(*conf); err != nil {
		log.Printf("[ERR]  err=%s\n", err)
	}

	//初始化配置文件
	if err := pkg.InitViperConf(); err != nil {
		log.Printf("[ERR]  err=%s\n", err)
	}

	// 初始化Logger
	if err := pkg.InitLogger(); err != nil {
		log.Printf("[ERR]  err=%s\n", err)
	}

	// 设置时区
	if _, err := time.LoadLocation(defined.TIME_LOCATION); err != nil {
		log.Printf("[ERR]  err=%s\n", err)
	}

	log.Printf("[INFO] %s\n", " success loading resources.")
	log.Println("------------------------------------------------------------------------")
}
