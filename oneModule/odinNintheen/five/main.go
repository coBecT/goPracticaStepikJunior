package main

import "fmt"

func main() {
	s := 0
	fmt.Scan(&s)
	if s <= 10000 && s > 0 {
		if s%400 == 0 || s%4 == 0 && s%100 != 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		fmt.Println()
	}
}

//fmt.Scan(&s)
//fmt.Println("YES")
//fmt.Println(s % 400)
//
//fmt.Println("NO")
//fmt.Println(s%4 == 0)
//fmt.Println(s%100 != 0)

//if s%400 == 0 {
//
//} else if s%4 == 0 && s%100 != 0 {
//
//}
