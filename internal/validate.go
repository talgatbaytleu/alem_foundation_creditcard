package internal

import (
	"creditcard/pkg"
)

func Validate() {
}

func LuhnCheck(s string) bool {
	if pkg.IsNumeric(s) && len(s) >= 12 && len(s) <= 16 {
		s = pkg.ReverseString(s)
		var res int32
		for i, v := range s {
			if i%2 != 0 {
				v = (v - 48) * 2
				if v < 10 {
					res += v
				} else {
					res += v - 9
				}
			} else {
				res += v - 48
			}
		}
		return res%10 == 0
	}
	return false
}
