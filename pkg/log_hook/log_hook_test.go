package log_hook

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogHook(t *testing.T) {
	h := New()

	logrus.AddHook(h)

	logrus.Info("hello")

}
