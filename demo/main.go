package main

import (
	"time"

	"github.com/yezi/mylog"
)

var (
	app  = "mylog"
	name = "test"
)

func main() {
	for {
		mylog.Logger.Infof("new mylog done,update status ok: app=%+v, name=%+v", app, name)
		mylog.Logger.Debugf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
		mylog.Logger.Warnf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
		mylog.Logger.Errorf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
		time.Sleep(time.Duration(3 * time.Second))
	}
}
