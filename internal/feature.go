package internal

import "os"

func InitFeature(feature_var string) {
	switch feature_var {
	case "validate":
		Validate()
	case "generate":
		Generate()
	case "information":
	case "issue":
	default:
		os.Exit(1)
	}
}
