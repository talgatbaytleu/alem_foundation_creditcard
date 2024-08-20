package pkg

// func IsAlpha(s string) bool {
// 	for _, r := range s {
// 		if !unicode.IsLetter(r) {
// 			return false
// 		}
// 	}
// 	return true
// }

func IsAlpha(s string) bool {
	for _, v := range s {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			return true
		}
	}
	return false
}
