package internal

import (
	"os"
)

var (
	Feature_var = InitFeatureVar()
	Flags_var   = os.Args[2:]
)

func InitFeatureVar() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	os.Exit(1)
	return ""
}
