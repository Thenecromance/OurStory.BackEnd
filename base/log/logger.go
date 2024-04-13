package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

//
//const (
//	LogPath   = "setting/stories.log"
//	ErrorPath = "setting/error.log"
//)
//
//var (
//	logger *zap.SugaredLogger
//	writer io.Writer
//)
//
//type Writer struct {
//}
//
//func (lw *Writer) Write(p []byte) (n int, err error) {
//	Get().Info(string(p))
//	return 0, nil
//}
//
//func init() {
//	loggers, err := setupConfig().Build(zap.ErrorOutput(zapcore.AddSync(&sqlWriter{})))
//	if err != nil {
//		panic(err)
//	}
//	logger = loggers.Sugar()
//	logger.Error("Test")
//}
//
//func setupEncoderConfig() zapcore.EncoderConfig {
//	return zapcore.EncoderConfig{
//		TimeKey:        "time",
//		LevelKey:       "level",
//		NameKey:        "log",
//		CallerKey:      "caller",
//		MessageKey:     "msg",
//		StacktraceKey:  "stacktrace",
//		LineEnding:     zapcore.DefaultLineEnding,
//		EncodeLevel:    zapcore.LowercaseLevelEncoder,
//		EncodeTime:     zapcore.ISO8601TimeEncoder,
//		EncodeDuration: zapcore.SecondsDurationEncoder,
//		EncodeCaller:   zapcore.ShortCallerEncoder,
//	}
//}
//
//func setupConfig() zap.Config {
//	return zap.Config{
//		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
//		Development:      false,
//		Encoding:         "json",
//		EncoderConfig:    setupEncoderConfig(),
//		OutputPaths:      []string{"stdout", LogPath},
//		ErrorOutputPaths: []string{"stderr", ErrorPath},
//	}
//
//}
//
//func Get() *zap.SugaredLogger {
//	file, _ := os.Create("./test.log")
//	writeSyncer := zapcore.AddSync(file)
//	core := zapcore.NewCore(zapcore.NewJSONEncoder(setupEncoderConfig()), writeSyncer, zapcore.DebugLevel)
//	test := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
//	test.Info("Test")
//	return logger
//}
//
//func GetWriter() io.Writer {
//	if writer == nil {
//		writer = &Writer{}
//	}
//	return writer
//
//}
//
//type logV2 struct {
//	logger *zap.SugaredLogger
//}
//
//func New() *logV2 {
//
//	return &logV2{}
//}

var (
	Instance *Log
)

const (
	logDir  = "log"
	logFile = logDir + "/" + "log.log"
	errFile = logDir + "/" + "err.log"

	skiplevel  = 2
	resetlevel = -2
)

type Log struct {
	logger  *zap.SugaredLogger
	writers *writerContainer
}

func (l *Log) initCore() {
	core := zapcore.NewCore(
		//zapcore.NewJSONEncoder(setupEncoderConfig()),
		zapcore.NewConsoleEncoder(setupEncoderConfig()),
		zapcore.AddSync(l.writers),
		zapcore.DebugLevel)

	l.logger = zap.New(core, zap.AddCaller()).Sugar()

	// let log output to the console
	l.writers.addWriter(os.Stderr)

	// let log output to the file
	{
		file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE|os.O_APPEND, os.ModeAppend)

		if err != nil {
			fmt.Println("open log file failed")
			return
		}

		l.writers.addWriter(file)
	}
}

// AppendWriter when log need to place to more place, just add the target to here
// just like if I want to log the data to the database, just need to wrap the sql
func (l *Log) AppendWriter(writer io.Writer) {
	l.writers.addWriter(writer)
}

func (l *Log) GetWriter() io.Writer {
	return l.writers
}

func (l *Log) WithOptions(opts ...zap.Option) *Log {
	l.logger.WithOptions(opts...)
	return l
}

func (l *Log) addCallerSkip() {
	l.logger = l.logger.WithOptions(zap.AddCallerSkip(skiplevel))
}
func (l *Log) resetCallerSkip() {
	l.logger = l.logger.WithOptions(zap.AddCallerSkip(resetlevel))
}

func New() *Log {
	l := &Log{
		writers: newWriterContainer(),
	}
	l.initCore()

	return l
}
