package internal

import (
	"fmt"
	"os"
)

// This file contains only one function, that call the corresponding function depending on Feature
// It parses only for corresponding flags
func InitFeature() {
	switch Feature_var {
	case "validate":
		ValidateCmd.Parse(os.Args[2:])
		ValidateFeature()
	case "generate":
		GenerateCmd.Parse(os.Args[2:])
		GenerateFeature()
	case "information":
		InformationCmd.Parse(os.Args[2:])
		InformationFeature()
	case "issue":
		IssueCmd.Parse(os.Args[2:])
		IssueFeature()
	default:
		fmt.Println("expected validate/generate/information/issue subcommands")
		os.Exit(1)
	}
}
