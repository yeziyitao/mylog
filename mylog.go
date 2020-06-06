package mylog

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-mesh/openlogging"
	mylog "github.com/yezi/mylog/loger"
)

// constant values for logrotate parameters
const (
	LogRotateDate     = 1
	LogRotateSize     = 10
	LogBackupCount    = 7
	RollingPolicySize = "size"
)

// log level
const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
	LevelFatal = "FATAL"
)

// output type
const (
	Stdout = "stdout"
	Stderr = "stderr"
	File   = "file"
)

//Logger is the global variable for the object of mylog.Logger
//Deprecated. plz use openlogging instead
var Logger mylog.Logger

// logFilePath log file path
var logFilePath string

//Options is the struct for mylog information(mylog.yaml)
type Options struct {
	Writers        string `yaml:"writers"`
	LoggerLevel    string `yaml:"logger_level"`
	LoggerFile     string `yaml:"logger_file"`
	LogFormatText  bool   `yaml:"log_format_text"`
	RollingPolicy  string `yaml:"rollingPolicy"`
	LogRotateDate  int    `yaml:"log_rotate_date"`
	LogRotateSize  int    `yaml:"log_rotate_size"`
	LogBackupCount int    `yaml:"log_backup_count"`

	AccessLogFile string `yaml:"access_log_file"`
}

// Init Build constructs a *Lager.Logger with the configured parameters.
func Init(option *Options) {
	var err error
	Logger, err = NewLog(option)
	if err != nil {
		panic(err)
	}
	openlogging.SetLogger(Logger)
	// openlogging.Debug("logger init success")
	return
}

func toLogLevel(option string) (mylog.LogLevel, error) {
	logLevel := mylog.DEBUG
	switch option {
	case LevelDebug:
	case LevelInfo:
		logLevel = mylog.INFO
	case LevelWarn:
		logLevel = mylog.WARN
	case LevelError:
		logLevel = mylog.ERROR
	case LevelFatal:
		logLevel = mylog.FATAL
	default:
		return 0, errors.New("invalid log level, valid: DEBUG, INFO, WARN, ERROR, FATAL")
	}

	return logLevel, nil
}

func toFile(writer string) (*os.File, error) {
	switch writer {
	case Stdout:
		return os.Stdout, nil
	case Stderr:
		return os.Stderr, nil
	case File:
		return os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	}
	return os.Stdout, nil
}

// NewLog returns a logger
func NewLog(option *Options) (mylog.Logger, error) {
	checkPassLagerDefinition(option)

	localPath := ""
	if !filepath.IsAbs(option.LoggerFile) {
		localPath = os.Getenv("CHASSIS_HOME")
	}
	err := createLogFile(localPath, option.LoggerFile)
	if err != nil {
		return nil, err
	}

	logFilePath = filepath.Join(localPath, option.LoggerFile)
	writers := strings.Split(strings.TrimSpace(option.Writers), ",")

	logger := mylog.NewLoggerExt(logFilePath, option.LogFormatText)
	option.LoggerFile = logFilePath

	logLevel, err := toLogLevel(option.LoggerLevel)
	if err != nil {
		return nil, err
	}

	for _, writer := range writers {
		f, err := toFile(writer)
		if err != nil {
			return nil, err
		}
		sink := mylog.NewReconfigurableSink(mylog.NewWriterSink(writer, f, mylog.DEBUG), logLevel)
		logger.RegisterSink(sink)
	}

	Rotators.Rotate(NewRotateConfig(option))
	return logger, nil
}

// checkPassLagerDefinition check pass mylog definition
func checkPassLagerDefinition(option *Options) {
	if option.LoggerLevel == "" {
		option.LoggerLevel = "DEBUG"
	}

	if option.LoggerFile == "" {
		option.LoggerFile = "log/mylog.log"
	}

	if option.RollingPolicy == "" {
		log.Println("RollingPolicy is empty, use default policy[size]")
		option.RollingPolicy = RollingPolicySize
	} else if option.RollingPolicy != "daily" && option.RollingPolicy != RollingPolicySize {
		log.Printf("RollingPolicy is error, RollingPolicy=%s, use default policy[size].", option.RollingPolicy)
		option.RollingPolicy = RollingPolicySize
	}

	if option.LogRotateDate <= 0 || option.LogRotateDate > 10 {
		option.LogRotateDate = LogRotateDate
	}

	if option.LogRotateSize <= 0 || option.LogRotateSize > 50 {
		option.LogRotateSize = LogRotateSize
	}

	if option.LogBackupCount < 0 || option.LogBackupCount > 100 {
		option.LogBackupCount = LogBackupCount
	}
}

// createLogFile create log file
func createLogFile(localPath, out string) error {
	_, err := os.Stat(strings.Replace(filepath.Dir(filepath.Join(localPath, out)), "\\", "/", -1))
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(strings.Replace(filepath.Dir(filepath.Join(localPath, out)), "\\", "/", -1), os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	f, err := os.OpenFile(strings.Replace(filepath.Join(localPath, out), "\\", "/", -1), os.O_CREATE, 0640)
	if err != nil {
		return err
	}
	return f.Close()
}
