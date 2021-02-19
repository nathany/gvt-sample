package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/nathany/gvt-sample/web"
)

func main() {
	// TODO(ny): do something awesome here
	logrus.Info("hello, universe")
	web.Start(":8080")
}
