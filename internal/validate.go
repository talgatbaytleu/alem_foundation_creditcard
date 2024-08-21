package internal

import (
	"fmt"
	"os"

	"creditcard/pkg"
)

func Validate() {
	if len(Args) == 1 && Args[0] == "--stdin" && len(Stdin_var) > 0 {
		for _, v := range Stdin_var {
			if LuhnCheck(v) {
				fmt.Println("OK")
			} else {
				fmt.Println("INCORRECT")
			}
		}
	} else if len(Args) > 0 {
		for _, v := range Args {
			if LuhnCheck(v) {
				fmt.Println("OK")
			} else {
				fmt.Fprintf(os.Stderr, "INCORRECT\n")
				os.Exit(1)
			}
		}
	} else {
		os.Exit(1)
	}
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
