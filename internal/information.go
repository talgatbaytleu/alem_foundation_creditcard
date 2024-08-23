package internal

import (
	"fmt"
	"os"
	"strings"
)

func InformationFeature() {
	// storing --brands and --issuers values to variables
	brands_file_name := *InformationBrands
	issuers_file_name := *InformationIssuers
	fmt.Println(*InformationStdin)

	// validating --brands and --issuers flags were specified
	if issuers_file_name == "" || brands_file_name == "" {
		fmt.Println("brands or issuers file not specified")
		os.Exit(1)
	}
	// Checking if the --stdin flag was enabled
	if len(InformationCmd.Args()) > 0 && !*InformationStdin {
		// if --stdin disabled
		for _, v := range InformationCmd.Args() {
			if LuhnCheck(v) {
				fmt.Println("Correct: yes")
				fmt.Println("Card Brand: " + VerifyBrandOrIssuer(v, brands_file_name))
				fmt.Println("Card Issuer: " + VerifyBrandOrIssuer(v, issuers_file_name))
			} else {
				fmt.Printf("Correct: no\nCard Brand: -\nCard Issuer: -\n")
			}
			fmt.Println("--------------------")
		}
	} else if len(InformationCmd.Args()) == 0 && Args_len == 3 && *InformationStdin {
		// if --stdin enabled
		stdin_var := InitStdinVar()
		for _, v := range stdin_var {
			if LuhnCheck(v) {
				fmt.Println("Correct: yes")
				fmt.Println("Card Brand: " + VerifyBrandOrIssuer(v, brands_file_name))
				fmt.Println("Card Issuer: " + VerifyBrandOrIssuer(v, issuers_file_name))
			} else {
				fmt.Printf("Correct: no\nCard Brand: -\nCard Issuer: -\n")
			}
			fmt.Println("--------------------")
		}
	} else {
		// Exiting if any error occurs
		fmt.Println("wrong usage of information subcommand")
		os.Exit(1)
	}
}

// Verifying Brand or Issuer name accorfing to provided BIN and file_name
func VerifyBrandOrIssuer(bin, file_name string) string {
	file, err := os.Open("../data/" + file_name)
	if err != nil {
		fmt.Println("some problem with specified files brands/issuers")
		os.Exit(1)
	}
	defer file.Close()

	file_content, _ := os.ReadFile(file.Name())
	file_content_str := string(file_content)

	lines := strings.Split(file_content_str, "\n")
	for _, v := range lines {
		parts := strings.Split(v, ":")
		if len(parts) == 2 && strings.HasPrefix(bin, parts[1]) {
			return parts[0]
		}
	}
	return "-"
}
