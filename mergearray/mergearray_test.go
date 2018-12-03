package mergearray

import "testing"

func TestMerge(t *testing.T) {
	tests := []struct {
		a    []int
		b    []int
		want []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{}, []int{1}, []int{1}},
		{[]int{1}, []int{1}, []int{1, 1}},
		{[]int{1}, []int{2}, []int{1, 2}},
		{[]int{2}, []int{1}, []int{1, 2}},
		{[]int{1, 3}, []int{0, 1, 4}, []int{0, 1, 1, 3, 4}},
	}
	for _, test := range tests {
		if got := Merge(test.a, test.b); !eq(got, test.want) {
			t.Errorf("Merge(%v, %v) = %v, want %v", test.a, test.b, got, test.want)
		}
	}
}

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
