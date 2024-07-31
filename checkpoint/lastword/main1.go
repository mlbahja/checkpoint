package main

import "fmt"

func main() {
	s := "test1 test2 tres3 "
	res := ""
	slic := []string{}
	for _, v := range s {
		if v == 'e' {
			slic = append(slic, res)
			res = ""
		} else {
			res += string(v)
		}
	}
	if res != "" {
		slic = append(slic, res)
	}
	fmt.Println(slic)
}
