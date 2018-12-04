package removeeven

// Remove removes even numbers from slice. It reuses underlying array.
func Remove(a []int) []int {
	odd := a[:0]
	for _, n := range a {
		if n%2 != 0 {
			odd = append(odd, n)
		}
	}
	return odd
}
