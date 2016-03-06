package loggomega

import "fmt"

type DefaultLogger struct{}

func (d *DefaultLogger) Infoln(args ...interface{}) {
	fmt.Println(args...)
}

func (d *DefaultLogger) Errorln(args ...interface{}) {
	fmt.Println(args...)
}
