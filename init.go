//Package mylog init necessary module
// before every other package init functions
package mylog

import (
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// LoggerOptions has the configuration about logging
var LoggerOptions *Options

func init() {
	InitLogger()
	// add Listener
	go addMyLogEventListener()
}

// InitLogger initiate config file and openlogging before other modules
func InitLogger() {
	err := ParseLoggerConfig("conf/mylog.yaml")
	//initialize log in any case
	if err != nil {
		Init(&Options{
			LoggerLevel:   LevelInfo,
			RollingPolicy: "size",
			Writers:       Stdout,
		})
		if os.IsNotExist(err) {
			Logger.Infof("[%s] not exist", "conf/mylog.yaml")
		} else {
			// log.Panicln(err)
			Logger.Errorf("config error: %s", err)
		}
	} else {
		Init(&Options{
			Writers:          LoggerOptions.Writers,
			LoggerLevel:      LoggerOptions.LoggerLevel,
			RollingPolicy:    LoggerOptions.RollingPolicy,
			LoggerFile:       LoggerOptions.LoggerFile,
			LogFormatText:    LoggerOptions.LogFormatText,
			LoggerFileFormat: LoggerOptions.LoggerFileFormat,
			LogRotateDate:    LoggerOptions.LogRotateDate,
			LogRotateSize:    LoggerOptions.LogRotateSize,
			LogBackupCount:   LoggerOptions.LogBackupCount,
			ZipOn:            LoggerOptions.ZipOn,
		})

	}
}

// ParseLoggerConfig unmarshals the logger configuration file(mylog.yaml)
func ParseLoggerConfig(file string) error {
	LoggerOptions = &Options{}
	err := unmarshalYamlFile(file, LoggerOptions)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return err
}

// unmarshalYamlFile read yaml conf file
func unmarshalYamlFile(file string, target interface{}) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(content, target)
}

// addMyLogEventListener watch mylog.yaml erver second
func addMyLogEventListener() {
	for {
		InitLogger()
		time.Sleep(time.Duration(1 * time.Second))
	}
}
