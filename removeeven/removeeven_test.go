package removeeven

import "testing"

func TestRemove(t *testing.T) {
	tests := []struct {
		a    []int
		want []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{2}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1}},
		{[]int{1, 2, 3, 4}, []int{1, 3}},
		{[]int{3, 4, 1, 2}, []int{3, 1}},
	}
	for _, test := range tests {
		if got := Remove(test.a); !eq(got, test.want) {
			t.Errorf("Remove(%v) = %v, want %v", test.a, got, test.want)
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
