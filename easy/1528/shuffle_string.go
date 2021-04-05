package easy

func restoreString(s string, indices []int) string {
	xi := make([]byte, len(indices))
	b := []byte(s)
	for i, pos := range indices {
		xi[pos] = b[i]
	}
	return string(xi)
}
