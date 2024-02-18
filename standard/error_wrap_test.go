package standard

import (
	"errors"
	"testing"
)

type errWrap struct {
	err error
}

func (e *errWrap) Error() string {
	return "普通 error 消息"
}

func (e *errWrap) Unwrap() error {
	return e.err
}

func (e *errWrap) Is(err error) bool {
	return e.Error() == err.Error()
}

func TestUnwrap(t *testing.T) {
	err1 := &errWrap{err: errors.New("封装的 error 消息")}
	t.Log("err:", err1)

	// Unwrap 解除 error 的封装，前提是 err1 实现了 Unwrap 方法，否则返回 nil
	ret := errors.Unwrap(err1)
	t.Log("unwrap error:", ret)
}

func TestIs(t *testing.T) {
	err1 := errors.New("ok")
	err2 := errors.New("ok")

	ret := errors.Is(err1, err2)
	t.Log("不同内存空间的 error 比较:", ret)

	err3 := errors.New("ok")
	err4 := err3

	ret = errors.Is(err3, err4)
	t.Log("相同空间的 error 比较", ret)

	err5 := &errWrap{err: errors.New("ok")}
	err6 := &errWrap{err: errors.New("ok")}
	ret = errors.Is(err5, err6)
	t.Log("不同内存空间的 error 用IS比较:", ret)
}

func TestAs(t *testing.T) {
	err1 := &errWrap{err: errors.New("ok")}

	var err2 = &errWrap{}

	errors.As(err1, &err2)
	t.Log(err1)
	t.Log(err2)
}
