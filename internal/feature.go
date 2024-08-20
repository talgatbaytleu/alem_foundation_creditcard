package internal

import "os"

func InitFeature(feature_var string) {
	switch feature_var {
	case "validate":
	case "generate":
	case "information":
	case "issue":
	default:
		os.Exit(1)
	}
}
