package internal

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Generating feature
func GenerateFeature() {
	var target_str string
	var asterisk_counter int
	if len(GenerateCmd.Args()) > 0 {
		// assigning las argument to target_str variable
		target_str = GenerateCmd.Args()[len(GenerateCmd.Args())-1]
		// Counting "*" (I believe there are a lot more better ways to do this)
		for strings.HasSuffix(target_str, "*") {
			asterisk_counter++
			target_str = target_str[:len(target_str)-1]
		}
		// validating number of * and length of target_str
		if asterisk_counter < 1 || asterisk_counter > 4 || len(target_str)+asterisk_counter < 12 ||
			len(target_str)+asterisk_counter > 16 {
			os.Exit(1)
		}
	}

	// generating slice which contains all posible varieties
	generated_slice := LuhnGenerate(target_str, asterisk_counter)

	// checking if the --pick flag enabled
	if len(GenerateCmd.Args()) == 1 && *GeneratePick {
		// if enabled
		randomNum := rand.Intn(len(generated_slice))
		fmt.Println(generated_slice[randomNum])
		os.Exit(0)
	} else if len(GenerateCmd.Args()) == 1 {
		// if disabled
		for _, v := range generated_slice {
			fmt.Println(v)
		}
		os.Exit(0)
	}
	os.Exit(1)
}

// It check all 10000(if we have 4 asterisks) via LuhnCheck function, then return slice of passed values
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
