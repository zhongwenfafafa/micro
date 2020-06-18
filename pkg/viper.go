package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/viper"

	"micro/defined"
)

var (
	ConfEnvPath  string
	ConfEnv      string
	ViperConfMap map[string]*viper.Viper
)

//初始化配置文件
func InitViperConf() error {
	basePath := ConfEnvPath + "/" + os.Getenv(defined.RUNTIME_ENV)
	f, err := os.Open(basePath + "/")
	if err != nil {
		return err
	}
	fileList, err := f.Readdir(1024)
	if err != nil {
		return err
	}
	for _, f0 := range fileList {
		if !f0.IsDir() {
			bts, err := ioutil.ReadFile(basePath + "/" + f0.Name())
			if err != nil {
				return err
			}
			v, err := ViperReadConfig(bts, "toml")
			if err != nil {
				return err
			}

			pathArr := strings.Split(f0.Name(), ".")
			if ViperConfMap == nil {
				ViperConfMap = make(map[string]*viper.Viper)
			}
			ViperConfMap[pathArr[0]] = v
		}
	}

	return nil
}

// 解析配置文件
func GetConfig(confName string, confStruct interface{}) error {
	var (
		ok bool
		v  *viper.Viper
	)
	ViperConfMap := ViperConfMap
	if v, ok = ViperConfMap[confName]; !ok {
		return fmt.Errorf("the %s not found in configuration items",
			confName)
	}

	if err := v.Unmarshal(confStruct); err != nil {
		return fmt.Errorf("parse config fail, config:%v, err:%v", confName, err)
	}

	return nil
}

func ParseConfPath(config string) error {
	path := strings.Split(config, "/")
	prefix := strings.Join(path[:len(path)-1], "/")
	ConfEnvPath = prefix
	ConfEnv = path[len(path)-2]

	return nil
}

//获取配置环境名
func GetConfEnv() string {
	return ConfEnv
}

func GetConfPath(fileName string) string {
	return ConfEnvPath + "/" + fileName + ".toml"
}

func GetConfFilePath(fileName string) string {
	return ConfEnvPath + "/" + fileName
}

// viper读配置文件
func ViperReadConfig(bts []byte, extensionName string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(extensionName)
	err := v.ReadConfig(bytes.NewBuffer(bts))

	if err != nil {
		return nil, err
	}

	return v, nil
}

//单独解析配置文件
func ParseConfig(filename string, conf interface{}) error {
	path := GetConfPath(filename)
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open config %v fail, %v", path, err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("read config fail, %v", err)
	}

	v, err := ViperReadConfig(data, "toml")
	if err != nil {
		return err
	}

	pathArr := strings.Split(file.Name(), ".")
	if ViperConfMap == nil {
		ViperConfMap = make(map[string]*viper.Viper)
	}
	ViperConfMap[pathArr[0]] = v

	if err := v.Unmarshal(conf); err != nil {
		return fmt.Errorf("parse config fail, config:%v, err:%v", string(data), err)
	}

	return nil
}
