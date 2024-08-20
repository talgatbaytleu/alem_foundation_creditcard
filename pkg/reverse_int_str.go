package pkg

func ReverseInt(x int) int {
	var res int
	for x != 0 {
		res *= 10
		res += x % 10
		x /= 10
	}
	return res
}

func ReverseString(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
