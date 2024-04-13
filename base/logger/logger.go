package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

const (
	LogPath   = "setting/stories.log"
	ErrorPath = "setting/error.log"
)

var (
	logger *zap.SugaredLogger
	writer io.Writer
)

//
//
//func GetWriter() io.Writer {
//	return sqlWriter
//}
//
//func init() {
//	Init()
//}
//
//func Init() {
//	sqlWriter = &loggerWriter{}
//
//	cfg := zap.NewProductionEncoderConfig()
//	//zap.NewDevelopmentConfig()
//	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
//
//	file, _ := os.Create(LogPath)
//
//	core := zapcore.NewTee(
//		zapcore.NewCore(zapcore.NewJSONEncoder(cfg), zapcore.AddSync(sqlWriter), zapcore.InfoLevel),
//		zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.AddSync(file), zapcore.InfoLevel),
//		zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
//	)
//	logger = zap.New(core).Sugar()
//
//}
//
//func AddCallback(callback func(p []byte) (n int, err error)) {
//	sqlWriter.callback = callback
//}
//func Get() *zap.SugaredLogger {
//	return logger
//}

type LoggerWriter struct {
}

func (lw *LoggerWriter) Write(p []byte) (n int, err error) {
	str := string(p)
	Get().Info(str)
	return 0, nil
}

func init() {
	loggers, err := setupConfig().Build()
	if err != nil {
		panic(err)
	}
	logger = loggers.Sugar()
}

func setupEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func setupConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    setupEncoderConfig(),
		OutputPaths:      []string{"stdout", LogPath},
		ErrorOutputPaths: []string{"stderr", ErrorPath},
	}
}

func Get() *zap.SugaredLogger {
	return logger
}

func GetWriter() io.Writer {
	if writer == nil {
		writer = &LoggerWriter{}
	}
	return writer
}
