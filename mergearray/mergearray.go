package mergearray

// Merge merges two sorted slices.
func Merge(a, b []int) []int {
	result := make([]int, len(a)+len(b))
	for i, j := 0, 0; i+j < len(result); {
		if j == len(b) {
			result[i+j] = a[i]
			i++
		} else if i == len(a) {
			result[i+j] = b[j]
			j++
		} else if a[i] <= b[j] {
			result[i+j] = a[i]
			i++
		} else {
			result[i+j] = b[j]
			j++
		}
	}
	return result
}
