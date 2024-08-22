package internal

import (
	"fmt"
	"os"
	"strings"
)

func VerifyBrandOrIssuer(bin, file_name string) string {
	file, err := os.Open("./data/" + file_name)
	if err != nil {
		fmt.Println("file with brands information not specified")
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
