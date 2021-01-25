package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
	// 不支持隐式类型转换，甚至同类型不同名的转换都不行
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	// 不支持指针运算
}

func TestSrting(t *testing.T) {
	var s string
	// 初始化为空字符串
	t.Log("*" + s + "*")
}
