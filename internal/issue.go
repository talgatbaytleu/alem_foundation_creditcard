package internal

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// This function for issuing 16 digit BIN based on brand and issuer
func IssueFeature() {
	// handling all file names and brand/issuer values according to flags
	brands_file_name := *IssueBrands
	issuers_file_name := *IssueIssuers
	chosen_brand := *IssueSpecifiedBrand
	chosen_issuer := *IssueSpecifiedIssuer
	// validating brands and issuers files were specified
	if issuers_file_name == "" || brands_file_name == "" {
		fmt.Println("brands or issuers file not specified")
		os.Exit(1)
	}
	// validating brand and issuer values were specified
	if chosen_brand == "" || chosen_issuer == "" {
		fmt.Println("brand or issuer not specified")
		os.Exit(1)
	}
	// validating brand and issuer were corresponding
	if len(IssueCmd.Args()) > 0 || Args_len > 4 {
		fmt.Println("Unknown flags")
		os.Exit(1)
	}
	if !strings.Contains(
		GetValuesOfBrandOrIssuer(issuers_file_name, chosen_issuer),
		GetValuesOfBrandOrIssuer(brands_file_name, chosen_brand),
	) {
		fmt.Println("Issuer does NOT match with Brand")
		os.Exit(1)
	}

	// Printing Generated BIN
	fmt.Println(GenerateToIssue(GetValuesOfBrandOrIssuer(issuers_file_name, chosen_issuer)))
	os.Exit(0)
}

// Function returns the corresponding value of the brand or issuer
func GetValuesOfBrandOrIssuer(file_name, brand_or_value string) string {
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
		if parts[0] == brand_or_value {
			return parts[1]
		}
	}
	return "none"
}

// Function generates BIN string based on brovided IIN(issuer identification number)
func GenerateToIssue(first_six_digits string) string {
	randomNum := rand.Intn(1000000000)
	first_six_digits += strconv.Itoa(randomNum)
	for i := 0; i < 15-len(first_six_digits); i++ {
		first_six_digits += "0"
	}
	return LuhnGenerate(first_six_digits, 1)[0]
}
