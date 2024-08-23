package internal

// This file contains all COMMON variables for whole program

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	// These all for subcommands with corresponding flags
	ValidateCmd   = flag.NewFlagSet("validate", flag.ExitOnError)
	ValidateStdin = ValidateCmd.Bool(
		"stdin",
		false,
		"creditcard validate --stdin to get input from standard input",
	)

	GenerateCmd  = flag.NewFlagSet("generate", flag.ExitOnError)
	GeneratePick = GenerateCmd.Bool(
		"pick",
		false,
		"creditcard generate --pick to pick randomly one BIN from generated list",
	)

	InformationCmd    = flag.NewFlagSet("information", flag.ExitOnError)
	InformationBrands = InformationCmd.String(
		"brands",
		"",
		"creditcard information --brands=fileName to chose file that stores brands' data (file must be stored in creditcard/data/ directory)",
	)
	InformationIssuers = InformationCmd.String(
		"issuers",
		"",
		"creditcard information --issuers=fileName to chose file that stores issuers' data (file must be stored in creditcard/data/ directory)",
	)
	InformationStdin = InformationCmd.Bool(
		"stdin",
		false,
		"creditcard information --stdin to get input from standard input",
	)

	IssueCmd    = flag.NewFlagSet("issue", flag.ExitOnError)
	IssueBrands = IssueCmd.String(
		"brands",
		"",
		"creditcard information --brands=fileName to chose file that stores brands' data (file must be stored in creditcard/data/ directory)",
	)
	IssueIssuers = IssueCmd.String(
		"issuers",
		"",
		"creditcard information --issuers=fileName to chose file that stores issuers' data (file must be stored in creditcard/data/ directory)",
	)
	IssueSpecifiedBrand = IssueCmd.String(
		"brand",
		"",
		"creditcard issue --brand=BrandName to chose what brand BIN should be issued",
	)
	IssueSpecifiedIssuer = IssueCmd.String(
		"issuer",
		"",
		"creditcard issue --issuer=IssuerName to chose what brand BIN should be issued",
	)

	// below variable just calls CheckFlags function
	check_for_equal_sign = CheckFlags()
	Feature_var          = InitFeatureVar()
	Args_len             = len(os.Args[2:])
)

// Function to check is there any argument, if so it will be assigned to Feature_var, whick stores subcomand value
func InitFeatureVar() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	fmt.Println("expected validate/generate/information/issue subcommands (InitFeatureVar)")
	os.Exit(1)
	return ""
}

// This function called only when --stdin flag used in order to avoid EOF
func InitStdinVar() []string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		words := strings.Fields(input)
		return words
	}
	return nil
}

// Function to validate flags format (NOT used)
func CheckFlags() bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && arg != "--stdin" && arg != "--pick" &&
			(!strings.Contains(arg, "=") ||
				!strings.HasPrefix(arg, "--")) {
			fmt.Println("Flags must be in the form --flag=vaule or --flag if it's boolian")
			os.Exit(1)
		}
	}
	return true
}
