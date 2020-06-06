## mylog简介

+ 支持修改日志级别实时生效,切割日志,删除历史日志
+ 开发时必要地方都加上Debug日志，线上排查问题，无需额外加日志排查，修改下日志级别即可

## 参考demo
```
mylog.Logger.Infof("new mylog done,update status ok: app=%+v, name=%+v", app, name)
mylog.Logger.Debugf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
mylog.Logger.Warnf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
mylog.Logger.Errorf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
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
