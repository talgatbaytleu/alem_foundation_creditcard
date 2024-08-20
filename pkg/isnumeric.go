package pkg

func IsNumeric(s string) bool {
	// if s[0] == '-' || s[0] == '+' {
	// 	for _, n := range s[1:] {
	// 		if 48 > n || n > 57 {
	// 			return false
	// 		}
	// 	}
	// } else {
	for _, n := range s {
		if 48 > n || n > 57 {
			return false
		}
	}
	// }

	return true
}
