package standard

import (
	"testing"
	"unsafe"
)

// SizeOf 返回的是这个对象的字节数
// slice, struct, 这种复杂结构的，返回的字节数不包含引用的内存部分
func TestSizeOf(t *testing.T) {
	var a int
	c := unsafe.Sizeof(a)
	t.Log("int", c)

	a2 := make([]int, 1000)
	c = unsafe.Sizeof(a2)
	t.Log("slice", c)

	var a3 = struct {
		a int
		b int
		d []int
	}{
		0, 1, make([]int, 1000),
	}
	c = unsafe.Sizeof(a3)
	t.Log("struct", c)

	m := make(map[string]string)
	c = unsafe.Sizeof(m)
	t.Log("map", c)

	ch := make(chan int, 100)
	c = unsafe.Sizeof(ch)
	t.Log("channel", c)
}

// Offsetof 返回的是 结构体 的字段偏移量
func TestOffsetOf(t *testing.T) {
	var a = struct {
		Name string
		Age  int
		Sex  int
	}{
		"Tom", 11, 1,
	}

	c := unsafe.Offsetof(a.Age)
	t.Log(c)

	c = unsafe.Offsetof(a.Sex)
	t.Log(c)

	p := unsafe.Pointer(&a)
	v := unsafe.Add(p, c)
	val := (*int)(v)
	t.Log(*val)
}

// Alignof 返回的是字段的对齐方式
func TestAlignof(t *testing.T) {
	var s struct {
		A bool
		B int16
		C []int
	}

	// s 在内存中的结构应该是
	// a 一个字节，补一个空字节，B占2个字节，补 4个空字节
	// C.data 占 8 字节
	// C.len 占 8 字节
	// C.cap 占 8 字节

	a := unsafe.Alignof(s.A)
	b := unsafe.Alignof(s.B)
	c := unsafe.Alignof(s.C)
	d := unsafe.Alignof(s)
	t.Logf("s: %d, A: %d, B: %d, C: %d", d, a, b, c)
}

func TestCompareSizeOfAndAlignOf(t *testing.T) {
	var s struct {
		A int
		B int
		C []int
	}
	r1 := unsafe.Sizeof(s)
	r2 := unsafe.Alignof(s)
	t.Logf("s size: %d, align: %d\n", r1, r2)

	r1 = unsafe.Sizeof(s.A)
	r2 = unsafe.Alignof(s.A)
	t.Logf("A size: %d, align: %d\n", r1, r2)

	r1 = unsafe.Sizeof(s.B)
	r2 = unsafe.Alignof(s.B)
	t.Logf("B size: %d, align: %d\n", r1, r2)

	r1 = unsafe.Sizeof(s.C)
	r2 = unsafe.Alignof(s.C)
	t.Logf("C size: %d, align: %d\n", r1, r2)
}

// unsafe.Pointer 计算
func TestAdd(t *testing.T) {
	a := []int{11, 12, 13}
	p := unsafe.Pointer(&a)

	v := *(*int)(unsafe.Add(p, 8))
	t.Log("取到的是 slice 结构体里的 cap", v)

	p1 := *(*unsafe.Pointer)(p) // p1 就是 slice 接口体里的 array unsafe.Pinter
	item := *(*int)(p1)
	t.Log("数组 index 1", item)

	item = *(*int)(unsafe.Add(p1, 8))
	t.Log("数组 index 2", item)
}

// Slice 返回一个新数组，需要传入数组其实位置的指针和数组长度
func TestSlice(t *testing.T) {
	a := make([]int, 10, 20)
	a[0] = 11
	a[2] = 12
	p := unsafe.Slice(&a[1], 10)
	t.Log(p)
	p2 := *(*[]int)(unsafe.Pointer(&p))
	t.Log(p2)
}

// ======= unsafe奇技淫巧 =======

// string to []byte 0 copy

func TestStringToBytes(t *testing.T) {
	a := "hello"
	pa := unsafe.Pointer(&a)

	b := *(*[]byte)(pa)
	t.Log(b)
	t.Log(string(b))
}

func TestUpdateFieldValue(t *testing.T) {
	a := struct {
		Name string
	}{
		"tom",
	}

	p := unsafe.Pointer(&a)

	namePtr := (*string)(p)
	*namePtr = "jack"
	t.Log(a.Name)
}
