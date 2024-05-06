package log

import "github.com/sirupsen/logrus"

func Init() {
	logrus.SetLevel(logrus.InfoLevel)
}
