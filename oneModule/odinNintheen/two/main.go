package main

import "fmt"

func main() {
	s := ""
	fmt.Scan(&s)
	trigger := true
	for i := 0; i < len(s); i++ {
		if s[0] == s[1] || s[1] == s[2] || s[2] == s[0] {
			trigger = false
		}
	}
	fmt.Println(string(s[0]))
	fmt.Println(string(s[1]))
	fmt.Println(string(s[2]))
	if trigger {
		fmt.Println("YES")

	} else {
		fmt.Println("NO")
	}
}
