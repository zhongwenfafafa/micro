package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"micro/defind"
	"micro/pkg"
)

type MysqlMapConf struct {
	Mysql map[string]*MySQLConf `mapstructure:"mysql"`
}

type MySQLConf struct {
	Driver          string `mapstructure:"driver"`
	Dsn             string `mapstructure:"dsn"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

var dbPool map[string]*gorm.DB

// 初始化数据库连接池
func InitDBPool() error {
	mysqlMap := &MysqlMapConf{}

	err := pkg.GetConfig("db", mysqlMap)
	if err != nil {
		return err
	}

	if len(mysqlMap.Mysql) == 0 {
		fmt.Printf("[INFO] %s%s\n", time.Now().Format(defind.TIME_FORMAT), " empty mysql config.")
	}

	dbPool = make(map[string]*gorm.DB)
	for dbType, dbConf := range mysqlMap.Mysql {
		dbGorm, err := gorm.Open("mysql", dbConf.Dsn)
		if err != nil {
			return err
		}

		dbGorm.SingularTable(true)
		err = dbGorm.DB().Ping()
		if err != nil {
			return err
		}

		dbGorm.LogMode(true)
		// 设置最大连接数
		dbGorm.DB().SetMaxOpenConns(dbConf.MaxOpenConn)
		// 设置最大空闲连接数
		dbGorm.DB().SetMaxIdleConns(dbConf.MaxIdleConn)
		// 设置连接超时最大时间
		dbGorm.DB().SetConnMaxLifetime(time.Duration(dbConf.MaxConnLifeTime))

		dbPool[dbType] = dbGorm
	}

	return nil
}

// 从db从库连接池中获取一个连接
func GetSlaveDBConn(ctx context.Context) (conn *gorm.DB, err error) {
	traceId := ctx.Value("traceId")
	if traceId == nil {
		return nil, errors.New("context not found traceId")
	}

	conn, err = pullConnInDBPool(defind.SLAVE_DB_NAME)
	if err != nil {
		return
	}

	conn.SetLogger(&pkg.GormLogger{TraceId: traceId.(string), Logger: pkg.Logger})
	return
}

// 从db主库连接池中获取一个连接
func GetMasterDBConn(ctx context.Context) (conn *gorm.DB, err error) {
	traceId := ctx.Value(defind.TRACE_KEY)
	if traceId == nil {
		return nil, errors.New("context not found traceId")
	}

	conn, err = pullConnInDBPool(defind.MASTER_DB_NAME)
	if err != nil {
		return
	}

	conn.SetLogger(&pkg.GormLogger{TraceId: traceId.(string), Logger: pkg.Logger})
	return
}

// 从连接池中获取连接
func pullConnInDBPool(dbType string) (*gorm.DB, error) {
	var (
		ok   bool
		conn *gorm.DB
	)
	if dbPool == nil {
		return nil, fmt.Errorf("mysql connection pool not init")
	}

	if conn, ok = dbPool[dbType]; !ok {
		return nil, fmt.Errorf("not found %s connect type in mysql connection pool",
			dbType)
	}

	return conn, nil
}
