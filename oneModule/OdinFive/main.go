package main

import "fmt"

func main() {
	res, s := 0.0, 0
	fmt.Scan(&s)
	res = float64(s % 10)
	fmt.Println(res)

}
