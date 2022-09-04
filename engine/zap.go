package engine

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var location, _ = time.LoadLocation("Asia/Shanghai")

type logCfg struct {
	Level      string
	FileName   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

func logInit(cfg *logCfg) (err error) {
	writeSyncer := getLogWriter(cfg.FileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	if err = l.UnmarshalText([]byte(cfg.Level)); err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	lg := zap.New(core)
	zap.ReplaceGlobals(lg)
	return
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = getDateTime
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getDateTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.In(location).Format("2006-01-02 15:04:05.999999"))
}

// 记录日志
func Log(msg string, err ...error) {
	if len(err) > 0 && err[0] != nil {
		zap.L().Error(msg, zap.Error(err[0]))
		fmt.Printf("%s => %v\n", msg, err)
	} else {
		zap.L().Info(msg)
		fmt.Printf("%s\n", msg)
	}
}
