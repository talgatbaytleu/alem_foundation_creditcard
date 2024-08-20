package pkg

// import "fmt"

// func main() {
// 	fmt.Println(Contains("hello world", "world"))
// 	fmt.Println(Contains("test", "best"))
// }

func Contains(str string, substr string) bool {
	if len(substr) > len(str) {
		return false
	}

	for i := range str[:len(str)-len(substr)+1] {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
