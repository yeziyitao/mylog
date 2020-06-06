package mylog

import (
	"fmt"
	"log"
	"testing"
	"time"
)

var (
	app  = "mylog"
	name = "test"
)

func TestMylog(t *testing.T) {
	fmt.Println("TestMylog")
	var i = 0
	for {
		i++
		if i == 3 {
			break
		}
		Logger.Infof("new mylog done,update status ok: app=%+v, name=%+v", app, name)
		Logger.Debugf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
		Logger.Warnf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
		Logger.Errorf("new mylog done,update status ok: app=%+v, name=%+v", app, name)
		log.Println("--")
		time.Sleep(time.Duration(3 * time.Second))
	}
}
