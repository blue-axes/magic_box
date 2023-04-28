package log_hook

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type (
	LogHook struct {
		skip int
	}
)

func New() *LogHook {
	h := &LogHook{
		skip: 4,
	}

	return h
}

func (l LogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l LogHook) Fire(entry *logrus.Entry) error {
	_, fn, line := getCaller(l.skip)
	entry.Data["func"] = fn
	entry.Data["line"] = line
	return nil
}

func getCaller(skip int) (string, string, int) {
	var (
		ptrs    = make([]uintptr, 10)
		n       = runtime.Callers(skip, ptrs)
		frams   = runtime.CallersFrames(ptrs[:n])
		hasNext = true
		f       runtime.Frame
	)
	for hasNext {
		f, hasNext = frams.Next()
		if strings.Contains(f.File, "logrus") {
			continue
		} else {
			return f.File, f.Function, f.Line
		}
	}
	return "", "", 0
}
