package leetcode

func reverseString(s string) string {
	b := []byte(s)
	start, end := 0, len(b)-1
	for end > start {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}

	return string(b)
}
