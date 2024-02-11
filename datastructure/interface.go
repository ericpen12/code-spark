package datastructure

import "unsafe"

type eface struct {
	_type *_type         // 实体的类型
	data  unsafe.Pointer // 指向实体的指针
}

type iface struct {
	tab  *itab          // 接口信息
	data unsafe.Pointer // 指向对象的地址，参考测试用例 Test_iface_data
}

type itab struct {
	inter *interfacetype // 接口信息
	_type *_type         // 实体的类型信息
	hash  uint32         // copy of _type.hash. Used for type switches.
	fun   [1]uintptr     // variable sized. fun[0]==0 means _type does not implement inter.
}

type interfacetype struct {
	typ     _type // 接口类型
	pkgpath name
	mhdr    []imethod
}
type name struct {
	bytes *byte
}

type imethod struct {
	name nameOff
	ityp typeOff
}

type nameOff int32
type typeOff int32
type tflag uint8

type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       nameOff // 函数名
	ptrToThis typeOff
}
