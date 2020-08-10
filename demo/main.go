package main

import (
	"time"

	"github.com/go-mesh/openlogging"
	"github.com/yeziyitao/mylog"
)

var (
	app  = "mylog"
	name = "test"
)

//LogFormat is a struct which stores details about log
type LogFormat struct {
	LogLevel  string `json:"level"`
	Timestamp string `json:"timestamp"`
	File      string `json:"file"`
	Message   string `json:"msg"`
}

func main() {
	format := LogFormat{
		LogLevel:  "INFO",
		Timestamp: "2020-07-31 16:24:22.635 +08:00",
		File:      "demo/main.go:36",
		Message:   "test",
	}
	for {
		// "msg":"new mylog done,update status ok: app=mylog, name=test"
		mylog.Infof("new mylog done,update status ok: app=%+v, name=%+v", app, name)
		mylog.Debugf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
		mylog.Warnf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
		mylog.Errorf("new mylog done,update status ok: app=%+v, name=%+v", app, name)

		// 打印结构体 "msg":"LogFormat:{LogLevel:DEBUG Timestamp:2020-07-31 16:24:22.635 +08:00 File:demo/main.go:36 Message:test}"
		mylog.Infof("LogFormat:%+v", format)
		// "msg":"LogFormat:{DEBUG 2020-07-31 16:24:22.635 +08:00 demo/main.go:36 test}"
		mylog.Infof("LogFormat:%s", format)

		time.Sleep(time.Duration(3000 * time.Millisecond))

		mylog.Info("start")
		// "msg":"begin","data":{"todo":"Start()"}
		mylog.Info("begin", openlogging.WithTags(openlogging.Tags{"todo": "Start()"}))
		// "msg":"begin","data":{"Address":"127.0.0.1:80","done":"ListenAndServe","protocol":"http"}
		mylog.Info("begin", openlogging.WithTags(openlogging.Tags{
			"protocol": "http",
			"done":     "ListenAndServe",
			"Address":  "127.0.0.1:80",
		}))
	}
}
