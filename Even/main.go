package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range s {
		if v%2 == 0 {
			fmt.Println(s[i], " is even")
		} else {
			fmt.Println(s[i], " is odd")
		}
	}
}
