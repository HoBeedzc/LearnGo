package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr4 := [2][2]int{{1, 2}, {3, 4}}
	arr1 = 5
	t.Log(arr[1], arr[2])
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		l.Log(arr3[i])
	}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[:] // 不支持负切片
	t.Log(arr3_sec)
}
