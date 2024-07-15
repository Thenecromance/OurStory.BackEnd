package log

import (
	"fmt"
	"github.com/Thenecromance/OurStories/constants"
	Config "github.com/Thenecromance/OurStories/utility/config"
	"github.com/Thenecromance/OurStories/utility/helper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// CustomZapConfig represents the YAML configuration structure for zap
type CustomZapConfig struct {
	Folder string `json:"folder" yaml:"folder"`
	// Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
	Level string `json:"level" yaml:"level"`
	// Development puts the logger in development mode, which changes the
	// behavior of DPanicLevel and takes stacktraces more liberally.
	Development bool `json:"development" yaml:"development"`
	// DisableCaller stops annotating logs with the calling function's file
	// name and line number. By default, all logs are annotated.
	DisableCaller bool `json:"disableCaller" yaml:"disableCaller"`
	// DisableStacktrace completely disables automatic stacktrace capturing. By
	// default, stacktraces are captured for WarnLevel and above logs in
	// development and ErrorLevel and above in production.
	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`
	// Encoding sets the logger's encoding. Valid values are "json" and
	// "console", as well as any third-party encodings registered via
	// RegisterEncoder.
	Encoding string `json:"encoding" yaml:"encoding"`
	// EncoderConfig sets options for the chosen encoder. See
	// zapcore.EncoderConfig for details.
	EncoderConfig EncoderConfig `json:"encoderConfig" yaml:"encoderConfig"`
	// OutputPaths is a list of URLs or file paths to write logging output to.
	// See Open for details.
	OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`
	// ErrorOutputPaths is a list of URLs to write internal logger errors to.
	// The default is standard error.
	//
	// Note that this setting only affects internal errors; for sample code that
	// sends error-level logs to a different location from info- and debug-level
	// logs, see the package-level AdvancedConfiguration example.
	ErrorOutputPaths []string `json:"errorOutputPaths" yaml:"errorOutputPaths"`
	// InitialFields is a collection of fields to add to the root logger.
	InitialFields map[string]interface{} `json:"initialFields" yaml:"initialFields"`
}

// EncoderConfig allows users to configure the concrete encoders supplied by zapcore.
type EncoderConfig struct {
	MessageKey       string `yaml:"messageKey"`
	LevelKey         string `yaml:"levelKey"`
	TimeKey          string `yaml:"timeKey"`
	NameKey          string `yaml:"nameKey"`
	CallerKey        string `yaml:"callerKey"`
	FunctionKey      string `yaml:"functionKey"`
	StacktraceKey    string `yaml:"stacktraceKey"`
	SkipLineEnding   bool   `yaml:"skipLineEnding"`
	LineEnding       string `yaml:"lineEnding"`
	EncodeLevel      string `yaml:"levelEncoder"`
	EncodeTime       string `yaml:"timeEncoder"`
	EncodeDuration   string `yaml:"durationEncoder"`
	EncodeCaller     string `yaml:"callerEncoder"`
	ConsoleSeparator string `yaml:"consoleSeparator"`

	//EncodeName       string `yaml:"nameEncoder"`
}

// LevelEncoderMapping holds the mapping between string values and zapcore.LevelEncoder functions.
var LevelEncoderMapping = map[string]zapcore.LevelEncoder{
	"capital":        zapcore.CapitalLevelEncoder,
	"lowercase":      zapcore.LowercaseLevelEncoder,
	"capitalColor":   zapcore.CapitalColorLevelEncoder,
	"lowercaseColor": zapcore.LowercaseColorLevelEncoder,
}
var encodeTimeMapping = map[string]func(time.Time, zapcore.PrimitiveArrayEncoder){
	"ISO8601":     zapcore.ISO8601TimeEncoder,
	"RFC3339":     zapcore.RFC3339TimeEncoder,
	"RFC3339Nano": zapcore.RFC3339NanoTimeEncoder,
}
var encodeDurationMapping = map[string]func(time.Duration, zapcore.PrimitiveArrayEncoder){
	"seconds": zapcore.SecondsDurationEncoder,
	"nanos":   zapcore.NanosDurationEncoder,
}
var encodeCallerMapping = map[string]func(zapcore.EntryCaller, zapcore.PrimitiveArrayEncoder){
	"short": zapcore.ShortCallerEncoder,
	"long":  zapcore.FullCallerEncoder,
}

// CustomLevelEncoder unmarshals a string to a zapcore.LevelEncoder.
func CustomLevelEncoder(name string) (zapcore.LevelEncoder, error) {
	if encoder, ok := LevelEncoderMapping[name]; ok {
		return encoder, nil
	}
	return nil, fmt.Errorf("unknown level encoder: %s", name)
}

func CustomTimeEncoder(name string) (func(time.Time, zapcore.PrimitiveArrayEncoder), error) {
	if encoder, ok := encodeTimeMapping[name]; ok {
		return encoder, nil
	}
	return nil, fmt.Errorf("unknown time encoder: %s", name)
}

func CustomDurationEncoder(name string) (func(time.Duration, zapcore.PrimitiveArrayEncoder), error) {
	if encoder, ok := encodeDurationMapping[name]; ok {
		return encoder, nil
	}
	return nil, fmt.Errorf("unknown duration encoder: %s", name)
}

func CustomCallerEncoder(name string) (func(zapcore.EntryCaller, zapcore.PrimitiveArrayEncoder), error) {
	if encoder, ok := encodeCallerMapping[name]; ok {
		return encoder, nil
	}
	return nil, fmt.Errorf("unknown caller encoder: %s", name)
}

// parseLevel converts a string level to a zapcore.Level
func parseLevel(level string) zapcore.Level {
	var l zapcore.Level
	if err := l.UnmarshalText([]byte(level)); err != nil {
		return zapcore.InfoLevel // default to Description if invalid level
	}
	return l
}

func defaultConfig() CustomZapConfig {
	return CustomZapConfig{
		Folder:            "logs",
		Level:             "debug",
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		InitialFields:     nil,
		Encoding:          "json",
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		EncoderConfig: EncoderConfig{
			MessageKey:       "msg",
			LevelKey:         "level",
			TimeKey:          "time",
			NameKey:          "logger",
			CallerKey:        "caller",
			FunctionKey:      zapcore.OmitKey,
			StacktraceKey:    "stacktrace",
			SkipLineEnding:   false,
			LineEnding:       "",
			EncodeLevel:      "lowercaseColor",
			EncodeTime:       "ISO8601",
			EncodeDuration:   "seconds",
			EncodeCaller:     "short",
			ConsoleSeparator: "",
		},
	}
}

func load() (*zap.Config, error) {
	customCfg := defaultConfig()
	Config.InstanceByName("zapConfig", constants.Yaml).LoadToObject("config", &customCfg)

	if customCfg.Folder != "" {
		if !helper.DirExists(customCfg.Folder) {
			err := os.Mkdir(customCfg.Folder, 0755)
			if err != nil {
				panic(err)
			}
		}
	}

	levelEncoder, err := CustomLevelEncoder(customCfg.EncoderConfig.EncodeLevel)
	if err != nil {
		levelEncoder = zapcore.LowercaseLevelEncoder
	}
	timeEncoder, err := CustomTimeEncoder(customCfg.EncoderConfig.EncodeTime)
	if err != nil {
		timeEncoder = zapcore.ISO8601TimeEncoder
	}
	durationEncoder, err := CustomDurationEncoder(customCfg.EncoderConfig.EncodeDuration)
	if err != nil {
		durationEncoder = zapcore.SecondsDurationEncoder
	}
	callerEncoder, err := CustomCallerEncoder(customCfg.EncoderConfig.EncodeCaller)
	if err != nil {
		callerEncoder = zapcore.ShortCallerEncoder
	}

	zapCfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(parseLevel(customCfg.Level)),
		Development:       customCfg.Development,
		DisableCaller:     customCfg.DisableCaller,
		DisableStacktrace: customCfg.DisableStacktrace,
		InitialFields:     customCfg.InitialFields,
		Encoding:          customCfg.Encoding,
		OutputPaths:       customCfg.OutputPaths,
		ErrorOutputPaths:  customCfg.ErrorOutputPaths,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     customCfg.EncoderConfig.MessageKey,
			LevelKey:       customCfg.EncoderConfig.LevelKey,
			TimeKey:        customCfg.EncoderConfig.TimeKey,
			NameKey:        customCfg.EncoderConfig.NameKey,
			CallerKey:      customCfg.EncoderConfig.CallerKey,
			FunctionKey:    customCfg.EncoderConfig.FunctionKey,
			StacktraceKey:  customCfg.EncoderConfig.StacktraceKey,
			SkipLineEnding: customCfg.EncoderConfig.SkipLineEnding,
			LineEnding:     customCfg.EncoderConfig.LineEnding + zapcore.DefaultLineEnding,
			EncodeLevel:    levelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: durationEncoder,
			EncodeCaller:   callerEncoder,

			// Add parsers for other encoders as needed
			ConsoleSeparator: customCfg.EncoderConfig.ConsoleSeparator,
		},
	}
	return &zapCfg, nil
}
