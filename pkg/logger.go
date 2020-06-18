package pkg

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"

	"micro/defined"

	"fmt"
	"os"
	"time"
)

type GormLogger struct {
	TraceId string
	Logger *zap.Logger
}

var Logger *zap.Logger

type LogMapConf struct {
	Log map[string]*LogConf `mapstructure:"log"`
}

type LogConf struct {
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

var logConfMap *LogMapConf

func InitLogger() error {
	env := os.Getenv(defined.RUNTIME_ENV)

	logConfMap = &LogMapConf{}
	err := GetConfig("log", logConfMap)
	if err != nil {
		return err
	}

	if len(logConfMap.Log) == 0 {
		fmt.Printf("[INFO] %s%s\n", time.Now().Format(defined.TIME_FORMAT), " empty log config.")
	}

	// 初始化日志
	// TODO 日志分类后期添加
	writeSyncer := logWriter(env)
	encoder := encoding()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	Logger = zap.New(core, zap.AddCaller())

	return nil
}

func (g *GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		g.Logger.Debug("execSql",
			zap.String("traceId", g.TraceId),
			zap.String("fileWithNum", v[1].(string)),
			zap.Duration("runtime", v[2].(time.Duration)),
			zap.String("sql", v[3].(string)),
			zap.Any("args", v[4]),
			zap.Any("affectNum", v[5]))

	case "log":
		g.Logger.Error("error",
			zap.String("traceId", g.TraceId),
			zap.String("fileWithNum", v[1].(string)),
			zap.Any("errorMessage", v[2]))

	case "info":
		g.Logger.Debug("log",
			zap.String("traceId", g.TraceId),
			zap.String("fileWithNum", v[1].(string)))
	}
}

// 日志文件编码
func encoding() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

// 日志文件写入和切割
func logWriter(prefix string) zapcore.WriteSyncer {
	sqlLogConf := logConfMap.Log["sql"]

	hook := &lumberjack.Logger{
		Filename:   "./runtime/" + strings.Trim(prefix, "/") + "/" + sqlLogConf.Filename,
		MaxSize:    sqlLogConf.MaxSize,
		MaxBackups: sqlLogConf.MaxBackups,
		MaxAge:     sqlLogConf.MaxSize,
		Compress:   sqlLogConf.Compress,
	}
	return zapcore.AddSync(hook)
}
