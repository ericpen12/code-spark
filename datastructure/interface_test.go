package datastructure

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"unsafe"
)

func Test_iface_data(t *testing.T) {
	buf := bytes.NewBufferString("Hello")
	var read io.Reader = buf
	i := (*iface)(unsafe.Pointer(&read))

	// i.data 指向的就是 buf
	b := (*bytes.Buffer)(i.data)
	b.WriteString(" world")

	t.Logf("b=%s, buf=%s", b, buf)
}

func Test_iFace_iTab_type(t *testing.T) {
	buf := bytes.NewBufferString("Hello")
	var read io.Reader = buf
	i := (*iface)(unsafe.Pointer(&read))

	tf := reflect.ValueOf(buf)

	t.Log(i.tab.inter.mhdr)

	t.Logf("%+v, %+v", i.tab, tf)
}
