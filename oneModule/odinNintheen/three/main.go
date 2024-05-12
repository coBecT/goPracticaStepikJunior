package main

import "fmt"

func main() {
	s, res := 0, 0
	fmt.Scan(&s)

	if s >= 10 && s < 100 {
		res = s / 10
	} else if s >= 100 && s < 1000 {
		res = s / 100
	} else if s >= 1000 && s < 10000 {
		res = s / 1000
	} else if s == 10000 {
		res = s / 10000
	} else if s < 10 {
		res = s
	}
	fmt.Println(res)
	//switch  {
	//	case :
	//
	//}
}
