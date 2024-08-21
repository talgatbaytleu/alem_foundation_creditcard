package internal

import (
	"bufio"
	"os"
	"strings"
)

var (
	Feature_var = InitFeatureVar()
	Args        = os.Args[2:]
	Stdin_var   = InitStdinVar()
)

func InitFeatureVar() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	os.Exit(1)
	return ""
}

func InitStdinVar() []string {
	if len(Args) == 1 && Args[0] == "--stdin" {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input := scanner.Text()
			words := strings.Fields(input)
			return words
		}
	}
	return nil
}
