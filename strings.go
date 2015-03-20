package strings

func Reverse(s string) string {
	chars := []rune(s)
	rev := make([]rune, len(chars))

	for i, j := len(chars)-1, 0; i >= 0; i, j = i-1, j+1 {
		rev[j] = chars[i]
	}

	return string(rev)
}
