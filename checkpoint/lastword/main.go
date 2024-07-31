package main

import "fmt"

func main() {
	str := "test1 test2 test3"
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
	fmt.Println(slice[len(slice)-1])
}
