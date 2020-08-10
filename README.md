## mylog简介

支持在线修改日志级别实时生效,切割日志,删除历史日志(开发时必要地方都加上Debug日志，线上排查问题，无需额外加日志排查，修改下日志级别即可)

+ 支持在线修改参数，不重启
+ 支持同时按照 文件大小 和 切割时间间隔 进行日志切割，支持秒级切割，文件大小按MB切割
+ 文件名路径 及 日志文件名格式支持自定义
+ 支持配置日志级别，日志可以输出到屏幕或者文件

## 代码参考demo/main.go，配置文件参考 demo/conf/mylog.yaml
```
        // {"level":"INFO","timestamp":"2020-08-05 10:30:31.921 +08:00","file":"demo/main.go:32","msg":"new mylog done,update status ok: app=mylog, name=test"}
        mylog.Infof("new mylog done,update status ok: app=%+v, name=%+v", app, name)
        mylog.Debugf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
        mylog.Warnf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
        mylog.Errorf("new mylog done,update status ok: app=%+v, name=%+v", app, name)

		// 打印结构体 "msg":"LogFormat:{LogLevel:DEBUG Timestamp:2020-07-31 16:24:22.635 +08:00 File:demo/main.go:36 Message:test}"
		mylog.Infof("LogFormat:%+v", format)
		// "msg":"LogFormat:{DEBUG 2020-07-31 16:24:22.635 +08:00 demo/main.go:36 test}"
		mylog.Infof("LogFormat:%s", format)

		mylog.Info("start")
		// "msg":"begin","data":{"todo":"Start()"}
		mylog.Info("begin", openlogging.WithTags(openlogging.Tags{"todo": "Start()"}))
		// "msg":"begin","data":{"Address":"127.0.0.1:80","done":"ListenAndServe","protocol":"http"}
		mylog.Info("begin", openlogging.WithTags(openlogging.Tags{
			"protocol": "http",
			"done":     "ListenAndServe",
			"Address":  "127.0.0.1:80",
		}))
```
## 示例,实时修改日志级别，不重启实现按照日志级别输出
Warn
```
MB1:demo yezi$ go run main.go 
2020/06/06 10:06:48 start
{"level":"WARN","timestamp":"2020-06-06 10:06:48.743 +08:00","file":"demo/main.go:22","msg":"new mylog done,update status ok: app=mylog, name=test"}
{"level":"ERROR","timestamp":"2020-06-06 10:06:48.744 +08:00","file":"demo/main.go:23","msg":"new mylog done,update status ok: app=mylog, name=test"}
2020/06/06 10:06:48 --
{"level":"WARN","timestamp":"2020-06-06 10:06:51.749 +08:00","file":"demo/main.go:22","msg":"new mylog done,update status ok: app=mylog, name=test"}
{"level":"ERROR","timestamp":"2020-06-06 10:06:51.749 +08:00","file":"demo/main.go:23","msg":"new mylog done,update status ok: app=mylog, name=test"}
```
info
```
{"level":"INFO","timestamp":"2020-06-06 10:07:06.769 +08:00","file":"demo/main.go:20","msg":"new mylog done,update status ok: app=mylog, name=test"}
{"level":"DEBUG","timestamp":"2020-06-06 10:07:06.770 +08:00","file":"demo/main.go:21","msg":"new mylog done,update status ok: app=mylog, name=test"}
{"level":"WARN","timestamp":"2020-06-06 10:07:06.770 +08:00","file":"demo/main.go:22","msg":"new mylog done,update status ok: app=mylog, name=test"}
{"level":"ERROR","timestamp":"2020-06-06 10:07:06.770 +08:00","file":"demo/main.go:23","msg":"new mylog done,update status ok: app=mylog, name=test"}
2020/06/06 10:07:06 --
{"level":"INFO","timestamp":"2020-06-06 10:07:09.773 +08:00","file":"demo/main.go:20","msg":"new mylog done,update status ok: app=mylog, name=test"}
{"level":"DEBUG","timestamp":"2020-06-06 10:07:09.773 +08:00","file":"demo/main.go:21","msg":"new mylog done,update status ok: app=mylog, name=test"}
{"level":"WARN","timestamp":"2020-06-06 10:07:09.773 +08:00","file":"demo/main.go:22","msg":"new mylog done,update status ok: app=mylog, name=test"}
{"level":"ERROR","timestamp":"2020-06-06 10:07:09.774 +08:00","file":"demo/main.go:23","msg":"new mylog done,update status ok: app=mylog, name=test"}
```
error
```
2020/06/06 10:07:09 --
{"level":"ERROR","timestamp":"2020-06-06 10:07:12.778 +08:00","file":"demo/main.go:23","msg":"new mylog done,update status ok: app=mylog, name=test"}
2020/06/06 10:07:12 --
{"level":"ERROR","timestamp":"2020-06-06 10:07:15.780 +08:00","file":"demo/main.go:23","msg":"new mylog done,update status ok: app=mylog, name=test"}
2020/06/06 10:07:15 --
```
warn
```
{"level":"WARN","timestamp":"2020-06-06 10:07:39.802 +08:00","file":"demo/main.go:22","msg":"new mylog done,update status ok: app=mylog, name=test"}
{"level":"ERROR","timestamp":"2020-06-06 10:07:39.803 +08:00","file":"demo/main.go:23","msg":"new mylog done,update status ok: app=mylog, name=test"}
2020/06/06 10:07:39 --
{"level":"WARN","timestamp":"2020-06-06 10:07:42.803 +08:00","file":"demo/main.go:22","msg":"new mylog done,update status ok: app=mylog, name=test"}
{"level":"ERROR","timestamp":"2020-06-06 10:07:42.804 +08:00","file":"demo/main.go:23","msg":"new mylog done,update status ok: app=mylog, name=test"}
```