package mylog

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-mesh/openlogging"
	"github.com/yeziyitao/mylog/loger"
)

// constant values for logrotate parameters
const (
	LogRotateDate     = 86400
	LogRotateSize     = 100
	LogBackupCount    = 100
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

//Logger is the global variable for the object of loger.Logger
//Deprecated. plz use openlogging instead
var Logger loger.Logger

// logFilePath log file path
var logFilePath string

//Options is the struct for loger information(loger.yaml)
type Options struct {
	Writers          string `yaml:"writers"`
	LoggerLevel      string `yaml:"logger_level"`
	LoggerFile       string `yaml:"logger_file"`
	LogFormatText    bool   `yaml:"log_format_text"`
	LoggerFileFormat string `yaml:"logger_file_format"`
	RollingPolicy    string `yaml:"rollingPolicy"`
	LogRotateDate    int    `yaml:"log_rotate_date"`
	LogRotateSize    int    `yaml:"log_rotate_size"`
	LogBackupCount   int    `yaml:"log_backup_count"`
	ZipOn            bool   `yaml:"zip_on"`
	AccessLogFile    string `yaml:"access_log_file"`
}

// Init Build constructs a *Lager.Logger with the configured parameters.
func Init(option *Options) {
	var err error
	Logger, err = NewLog(option)
	if err != nil {
		openlogging.Error("invalid log level:")
	}
	openlogging.SetLogger(Logger)
	// openlogging.Debug("logger init success")
	return
}

func toLogLevel(option string) (loger.LogLevel, error) {
	logLevel := loger.DEBUG
	switch option {
	case LevelDebug:
	case LevelInfo:
		logLevel = loger.INFO
	case LevelWarn:
		logLevel = loger.WARN
	case LevelError:
		logLevel = loger.ERROR
	case LevelFatal:
		logLevel = loger.FATAL
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
func NewLog(option *Options) (loger.Logger, error) {
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

	logger := loger.NewLoggerExt(logFilePath, option.LogFormatText)
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
		sink := loger.NewReconfigurableSink(loger.NewWriterSink(writer, f, loger.DEBUG), logLevel)
		logger.RegisterSink(sink)
	}
	// fmt.Println("NewRotateConfig.option:", option)
	Rotators.Rotate(NewRotateConfig(option))

	// fmt.Println("--logger--", logger)
	return logger, nil
}

// checkPassLagerDefinition check pass loger definition
func checkPassLagerDefinition(option *Options) {
	// fmt.Println("checkPassLagerDefinition>>>>>>>>>>>>>>>>>>>>>>:", option)
	// if option.LoggerLevel == "" || option.LoggerLevel != "DEBUG" || option.LoggerLevel != "INFO" || option.LoggerLevel != "WARN" || option.LoggerLevel != "ERROR" || option.LoggerLevel != "FATAL" {
	if option.LoggerLevel == "" {
		option.LoggerLevel = "INFO"
	}
	_, err := toLogLevel(option.LoggerLevel)
	if err != nil {
		openlogging.Error("invalid log level: " + option.LoggerLevel + ", use defalut INFO")
		option.LoggerLevel = "INFO"
	}

	if option.LoggerFile == "" {
		option.LoggerFile = "log/mylog.log"
	}

	if option.LoggerFileFormat == "" {
		option.LoggerFileFormat = "2006.01.02_15.04.05"
	}

	if option.LogRotateDate <= 0 {
		option.LogRotateDate = LogRotateDate
	}

	if option.LogRotateSize <= 0 {
		option.LogRotateSize = LogRotateSize
	}

	if option.LogBackupCount < 0 {
		option.LogBackupCount = LogBackupCount
	}

	// fmt.Println("option:", option)
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
