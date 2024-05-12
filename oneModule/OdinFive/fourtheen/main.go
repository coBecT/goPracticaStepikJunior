package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := 0
	fmt.Scan(&s)
	fmt.Println(s)
	if s < 10000 {
		a := strconv.Itoa(s)
		b := a[len(a)-2 : len(a)-1]
		fmt.Println(b)
	} else {
		fmt.Print("mnogo")
	}
}
