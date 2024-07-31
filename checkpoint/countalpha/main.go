package main

import "fmt"

func CouuntAlpha(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'z') {
			count++
		}
	}
	return count
}

func main() {
	num := CouuntAlpha("H1e2l3l4o")
	fmt.Println(num)
}
