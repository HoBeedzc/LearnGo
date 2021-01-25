package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(a == c) // 长度不同的数组比较时会出现编译错误
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	var a int
	a = 7             // 00111
	a = a &^ Readable // 按位清零
	t.Log(Monday, Tuesday)
}
