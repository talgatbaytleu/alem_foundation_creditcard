package internal

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func Generate() {
	var target_str string
	var asterisk_counter int
	if len(Args) > 0 {
		target_str = Args[len(Args)-1]
		for strings.HasSuffix(target_str, "*") {
			asterisk_counter++
			target_str = target_str[:len(target_str)-1]
		}
		if asterisk_counter < 1 || asterisk_counter > 4 || len(target_str)+asterisk_counter < 12 ||
			len(target_str)+asterisk_counter > 16 {
			os.Exit(1)
		}
	}

	generated_slice := LuhnGenerate(target_str, asterisk_counter)

	if len(Args) == 2 && Args[0] == "--pick" {
		randomNum := rand.Intn(len(generated_slice))
		fmt.Println(generated_slice[randomNum])
		os.Exit(0)
	} else if len(Args) == 1 {
		for _, v := range generated_slice {
			fmt.Println(v)
		}
		os.Exit(0)
	}
	os.Exit(1)
}

func LuhnGenerate(s string, asterisk_count int) []string {
	var res string
	n := int(math.Pow10(asterisk_count))
	res_slice := make([]string, 0, n)
	for i := 0; i < n; i++ {
		res = ""
		for j := 0; j < asterisk_count-(len(strconv.Itoa(i))); j++ {
			res += "0"
		}
		res += strconv.Itoa(i)
		if LuhnCheck(s + res) {
			res_slice = append(res_slice, s+res)
		}
	}
	return res_slice
}
