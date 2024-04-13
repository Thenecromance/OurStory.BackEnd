package log

import (
	"go.uber.org/zap/zapcore"
)

// Debug logs the provided arguments at [DebugLevel].
// Spaces are added between arguments when neither is a string.
func (l *Log) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

// Info logs the provided arguments at [InfoLevel].
// Spaces are added between arguments when neither is a string.
func (l *Log) Info(args ...interface{}) {
	l.logger.Info(args...)
}

// Warn logs the provided arguments at [WarnLevel].
// Spaces are added between arguments when neither is a string.
func (l *Log) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

// Error logs the provided arguments at [ErrorLevel].
// Spaces are added between arguments when neither is a string.
func (l *Log) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// DPanic logs the provided arguments at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
// Spaces are added between arguments when neither is a string.
func (l *Log) DPanic(args ...interface{}) {
	l.logger.DPanic(args...)
}

// Panic constructs a message with the provided arguments and panics.
// Spaces are added between arguments when neither is a string.
func (l *Log) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

// Fatal constructs a message with the provided arguments and calls os.Exit.
// Spaces are added between arguments when neither is a string.
func (l *Log) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Logf formats the message according to the format specifier
// and logs it at provided level.
func (l *Log) Logf(lvl zapcore.Level, template string, args ...interface{}) {
	l.logger.Logf(lvl, template, args...)
}

// Debugf formats the message according to the format specifier
// and logs it at [DebugLevel].
func (l *Log) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

// Infof formats the message according to the format specifier
// and logs it at [InfoLevel].
func (l *Log) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

// Warnf formats the message according to the format specifier
// and logs it at [WarnLevel].
func (l *Log) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

// Errorf formats the message according to the format specifier
// and logs it at [ErrorLevel].
func (l *Log) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

// DPanicf formats the message according to the format specifier
// and logs it at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
func (l *Log) DPanicf(template string, args ...interface{}) {
	l.logger.DPanicf(template, args...)
}

// Panicf formats the message according to the format specifier
// and panics.
func (l *Log) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

// Fatalf formats the message according to the format specifier
// and calls os.Exit.
func (l *Log) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

// Logw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Log) Logw(lvl zapcore.Level, msg string, keysAndValues ...interface{}) {
	l.logger.Logw(lvl, msg, keysAndValues...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func (l *Log) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Log) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Log) Warnw(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Log) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (l *Log) DPanicw(msg string, keysAndValues ...interface{}) {
	l.logger.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (l *Log) Panicw(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (l *Log) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

// Logln logs a message at provided level.
// Spaces are always added between arguments.
func (l *Log) Logln(lvl zapcore.Level, args ...interface{}) {
	l.logger.Logln(lvl, args...)
}

// Debugln logs a message at [DebugLevel].
// Spaces are always added between arguments.
func (l *Log) Debugln(args ...interface{}) {
	l.logger.Debugln(args...)
}

// Infoln logs a message at [InfoLevel].
// Spaces are always added between arguments.
func (l *Log) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

// Warnln logs a message at [WarnLevel].
// Spaces are always added between arguments.
func (l *Log) Warnln(args ...interface{}) {
	l.logger.Warnln(args...)
}

// Errorln logs a message at [ErrorLevel].
// Spaces are always added between arguments.
func (l *Log) Errorln(args ...interface{}) {
	l.logger.Errorln(args...)
}

// DPanicln logs a message at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
// Spaces are always added between arguments.
func (l *Log) DPanicln(args ...interface{}) {
	l.logger.DPanicln(args...)
}

// Panicln logs a message at [PanicLevel] and panics.
// Spaces are always added between arguments.
func (l *Log) Panicln(args ...interface{}) {
	l.logger.Panicln(args...)
}

// Fatalln logs a message at [FatalLevel] and calls os.Exit.
// Spaces are always added between arguments.
func (l *Log) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}

//------------------------------------------------------------

// Debug logs the provided arguments at [DebugLevel].
// Spaces are added between arguments when neither is a string.
func Debug(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()

	Instance.Debug(args...)

}

// Info logs the provided arguments at [InfoLevel].
// Spaces are added between arguments when neither is a string.
func Info(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()

	Instance.Info(args...)
}

// Warn logs the provided arguments at [WarnLevel].
// Spaces are added between arguments when neither is a string.
func Warn(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Warn(args...)
}

// Error logs the provided arguments at [ErrorLevel].
// Spaces are added between arguments when neither is a string.
func Error(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Error(args...)
}

// DPanic logs the provided arguments at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
// Spaces are added between arguments when neither is a string.
func DPanic(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.DPanic(args...)
}

// Panic constructs a message with the provided arguments and panics.
// Spaces are added between arguments when neither is a string.
func Panic(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Panic(args...)
}

// Fatal constructs a message with the provided arguments and calls os.Exit.
// Spaces are added between arguments when neither is a string.
func Fatal(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Fatal(args...)
}

// Logf formats the message according to the format specifier
// and logs it at provided level.
func Logf(lvl zapcore.Level, template string, args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Logf(lvl, template, args...)
}

// Debugf formats the message according to the format specifier
// and logs it at [DebugLevel].
func Debugf(template string, args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Debugf(template, args...)
}

// Infof formats the message according to the format specifier
// and logs it at [InfoLevel].
func Infof(template string, args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Infof(template, args...)
}

// Warnf formats the message according to the format specifier
// and logs it at [WarnLevel].
func Warnf(template string, args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Warnf(template, args...)
}

// Errorf formats the message according to the format specifier
// and logs it at [ErrorLevel].
func Errorf(template string, args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Errorf(template, args...)
}

// DPanicf formats the message according to the format specifier
// and logs it at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
func DPanicf(template string, args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.DPanicf(template, args...)
}

// Panicf formats the message according to the format specifier
// and panics.
func Panicf(template string, args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Panicf(template, args...)
}

// Fatalf formats the message according to the format specifier
// and calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Fatalf(template, args...)
}

// Logw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Logw(lvl zapcore.Level, msg string, keysAndValues ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Logw(lvl, msg, keysAndValues...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Fatalw(msg, keysAndValues...)
}

// Logln logs a message at provided level.
// Spaces are always added between arguments.
func Logln(lvl zapcore.Level, args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Logln(lvl, args...)
}

// Debugln logs a message at [DebugLevel].
// Spaces are always added between arguments.
func Debugln(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Debugln(args...)
}

// Infoln logs a message at [InfoLevel].
// Spaces are always added between arguments.
func Infoln(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Infoln(args...)
}

// Warnln logs a message at [WarnLevel].
// Spaces are always added between arguments.
func Warnln(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Warnln(args...)
}

// Errorln logs a message at [ErrorLevel].
// Spaces are always added between arguments.
func Errorln(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Errorln(args...)
}

// DPanicln logs a message at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
// Spaces are always added between arguments.
func DPanicln(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.DPanicln(args...)
}

// Panicln logs a message at [PanicLevel] and panics.
// Spaces are always added between arguments.
func Panicln(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Panicln(args...)
}

// Fatalln logs a message at [FatalLevel] and calls os.Exit.
// Spaces are always added between arguments.
func Fatalln(args ...interface{}) {
	Instance.addCallerSkip()
	defer Instance.resetCallerSkip()
	Instance.Fatalln(args...)
}
