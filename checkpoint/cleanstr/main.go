package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	if len(os.Args) != 1 {
		printstr("\n")
		return
	}
	str := os.Args[1]
	res := ""
	slice := []string{}

	for _, char := range str {
		if char == ' ' {
			slice = append(slice, res)
			res = ""
		} else {
			res += string(char)
		}
	}
	if res != "" {
		slice = append(slice, res)
	}
	for i, char := range slice {
		printstr(char)
		if len(slice)-1 > i {
			printstr(" ")
		}

	}

	// fmt.Println(slice)
}
