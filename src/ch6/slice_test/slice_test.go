package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int // slice 可变长
	t.Log(len(s0), cap(s0))
	s0 := append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(lem(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[4]) // len初始化的个数，cap代表容量
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s)) //cap增长是直接翻倍，log增长是一个一个涨
	}
}
