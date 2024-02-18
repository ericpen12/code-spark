package standard

import (
	"errors"
	"testing"
)

type customError struct {
}

func (c *customError) Error() string {
	return "custom error"
}

func TestError(t *testing.T) {
	var err error = &customError{}
	t.Log(err)
}

// TestNewEqual 错误比较
func TestNewEqual(t *testing.T) {
	// 内存不同的 error 不能比较.
	if errors.New("abc") == errors.New("abc") {
		t.Errorf(`New("abc") == New("abc")`)
	}
	if errors.New("abc") == errors.New("xyz") {
		t.Errorf(`New("abc") == New("xyz")`)
	}
	// Same allocation should be equal to itself (not crash).
	err := errors.New("jkl")
	if err != err {
		t.Errorf(`err != err`)
	}
}

func TestErrorMethod(t *testing.T) {
	err := errors.New("abc")
	if err.Error() != "abc" {
		t.Errorf(`New("abc").Error() = %q, want %q`, err.Error(), "abc")
	}
}
