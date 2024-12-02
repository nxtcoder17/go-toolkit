package log_test

import (
	"os"
	"testing"

	"github.com/nxtcoder17/go-toolkit/log"
)

func TestLoggerJson(t *testing.T) {
	l := log.New(log.Options{
		Prefix:          "testing",
		Writer:          os.Stderr,
		Theme:           log.ThemeDark,
		ShowCaller:      true,
		DevelopmentMode: false,
	})

	l.Debug("Hello World", "k1", "v1", "k2", 2, "k3", true, "k4", 2.15)
	l.Info("Hello World", "k1", "v1", "k2", 2, "k3", true, "k4", 2.15)
	l.Warn("Hello World", "k1", "v1", "k2", 2, "k3", true, "k4", 2.15)
	l.Error("Hello World", "k1", "v1", "k2", 2, "k3", true, "k4", 2.15)
}

func TestLoggerConsole(t *testing.T) {
	l := log.New(log.Options{
		Prefix:          "testing",
		Writer:          os.Stderr,
		Theme:           log.ThemeDark,
		ShowCaller:      true,
		DevelopmentMode: true,
	})

	l.Debug("Hello World", "k1", "v1", "k2", 2, "k3", true, "k4", 2.15)
	l.Info("Hello World", "k1", "v1", "k2", 2, "k3", true, "k4", 2.15)
	l.Warn("Hello World", "k1", "v1", "k2", 2, "k3", true, "k4", 2.15)
	l.Error("Hello World", "k1", "v1", "k2", 2, "k3", true, "k4", 2.15)
}
