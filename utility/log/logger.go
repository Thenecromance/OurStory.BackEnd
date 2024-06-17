package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

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
