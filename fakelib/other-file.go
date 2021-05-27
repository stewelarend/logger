package fakelib

import (
	"time"

	"github.com/stewelarend/logger"
)

func otherFileFunc(myLog logger.ILogger) {
	myLog.Debugf("Fake on debug")
	time.Sleep(time.Millisecond * 100)
	myLog = myLog.With("valid", true)
	myLog.Infof("Fake on debug")
	time.Sleep(time.Millisecond * 100)
	myLog.Error("Fake on debug")
	time.Sleep(time.Millisecond * 100)
}
