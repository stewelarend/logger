package logger_test

import (
	"fmt"
	"testing"

	"github.com/stewelarend/logger"
)

func TestError1(t *testing.T) {
	e1 := fmt.Errorf("error1")
	e2 := fmt.Errorf("error2: %v", e1)
	e3 := logger.Wrapf(e2, "error3")
	e4 := logger.Wrapf(e3, "error4")
	e5 := logger.Wrapf(e4, "error5")
	if s := fmt.Sprintf("%e", e5); s != "error2: error1" {
		t.Fatalf("%%e not correctly formatted: %s", s)
	}
	if s := fmt.Sprintf("%+e", e5); s != "error2: error1" {
		t.Fatalf("%%+e not correctly formatted: %s", s)
	}
	if s := fmt.Sprintf("%s", e5); s != "error5" {
		t.Fatalf("%%s not correctly formatted: %s", s)
	}
	if s := fmt.Sprintf("%+s", e5); s != "github.com/stewelarend/logger_test/errors_test.go(15): error5" {
		t.Fatalf("%%+s not correctly formatted: %s", s)
	}
	if s := fmt.Sprintf("%v", e5); s != "error5; error4; error3; error2: error1" {
		t.Fatalf("%%v not correctly formatted: %s", s)
	}
	if s := fmt.Sprintf("%+v", e5); s != "github.com/stewelarend/logger_test/errors_test.go(15): error5; github.com/stewelarend/logger_test/errors_test.go(14): error4; github.com/stewelarend/logger_test/errors_test.go(13): error3; error2: error1" {
		t.Fatalf("%%+v not correctly formatted: %s", s)
	}
}
