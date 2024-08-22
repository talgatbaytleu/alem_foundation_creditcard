package main

import (
	"fmt"

	"creditcard/internal"
)

func main() {
	// internal.InitFeature(internal.Feature_var)
	a := internal.VerifyBrandOrIssuer("4400430180300003", "brands.txt")
	fmt.Println(a)
	b := internal.VerifyBrandOrIssuer("4400430180300003", "issuers.txt")
	fmt.Println(b)
	// fmt.Println(internal.Feature_var)
	// fmt.Println(internal.Args)
	// fmt.Println(len(internal.Args))
	// fmt.Println(internal.LuhnCheck("4400430144050652"))
	// fmt.Println(internal.LuhnCheck("5269940015288691"))
	// fmt.Println(internal.LuhnCheck("4400430144050651"))
	// fmt.Println(internal.LuhnCheck("5269940015288692"))
	// fmt.Println(internal.Stdin_var)
	// fmt.Println(len(internal.Stdin_var))
}
