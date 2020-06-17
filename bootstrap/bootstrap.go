package bootstrap

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"micro/db"
	"micro/defind"
	"micro/pkg"
)

func InitModule(confPath string) error {
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
		return err
	}

	//初始化配置文件
	if err := pkg.InitViperConf(); err != nil {
		return err
	}

	// 加载mysql配置并初始化实例
	if err := db.InitDBPool(); err != nil {
		fmt.Printf("[ERROR] %s%s\n", time.Now().Format(defind.TIME_FORMAT), " InitDBPool:"+err.Error())
		os.Exit(1)
	}

	// 加载logger配置并初始化实例
	if err := pkg.InitLogger(); err != nil {
		fmt.Printf("[ERROR] %s%s\n", time.Now().Format(defind.TIME_FORMAT), " InitLogger:"+err.Error())
		os.Exit(1)
	}

	// 设置时区
	if _, err := time.LoadLocation(defind.TIME_LOCATION); err != nil {
		return err
	}

	log.Printf("[INFO] %s\n", " success loading resources.")
	log.Println("------------------------------------------------------------------------")
	return nil
}
