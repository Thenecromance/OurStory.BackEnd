package log

import (
	"go.uber.org/zap"
	"io"
)

var (
	Instance *Logger
)

const (
	logDir  = "logs"
	logFile = logDir + "/" + "server.log"
	errFile = logDir + "/" + "err_server.log"

	skiplevel  = 2
	resetlevel = -2
)

type Logger struct {
	_logger *zap.SugaredLogger
	writers *writerContainer
}

/*func (l *Logger) initCore() {

	cfg, _ := load()

	core := zapcore.NewCore(
		//zapcore.NewJSONEncoder(setupEncoderConfig()),
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(l.writers),
		zapcore.DebugLevel)

	l._logger = zap.New(core, zap.AddCaller()).Sugar()

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
*/

func (l *Logger) initCore() {
	cfg, err := load()
	build, err := cfg.Build()
	if err != nil {
		return
	}

	l._logger = build.Sugar()

}

// AppendWriter when log need to place to more place, just add the target to here
// just like if I want to log the data to the database, just need to wrap the sql
func (l *Logger) AppendWriter(writer io.Writer) {
	l.writers.addWriter(writer)
}

func (l *Logger) GetWriter() io.Writer {
	return l.writers
}

func (l *Logger) WithOptions(opts ...zap.Option) *Logger {
	l._logger.WithOptions(opts...)
	return l
}

func (l *Logger) addCallerSkip() {
	l._logger = l._logger.WithOptions(zap.AddCallerSkip(skiplevel))
}
func (l *Logger) resetCallerSkip() {
	l._logger = l._logger.WithOptions(zap.AddCallerSkip(resetlevel))
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l._logger.Debugf(format, v...)
}

func New() *Logger {
	l := &Logger{
		writers: newWriterContainer(),
	}
	l.initCore()

	return l
}
