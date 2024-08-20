package main

import (
	"fmt"

	"creditcard/internal"
)

func main() {
	internal.InitFeature(internal.Feature_var)
	fmt.Println(internal.Feature_var)
	fmt.Println(internal.Flags_var)
	fmt.Println(len(internal.Flags_var))
	fmt.Println(internal.LuhnCheck("4400430144050652"))
	fmt.Println(internal.LuhnCheck("5269940015288691"))
	fmt.Println(internal.LuhnCheck("4400430144050651"))
	fmt.Println(internal.LuhnCheck("5269940015288692"))
}
