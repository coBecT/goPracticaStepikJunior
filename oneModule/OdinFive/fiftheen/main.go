package main

import "fmt"

func main() {
	a, hour, min := 0, 0, 0
	fmt.Scan(&a)
	for a > 0 {
		if a >= 30 {
			hour++
			a -= 30
		} else if a < 30 {
			min = a + a
			a = 0
		}
	}

	fmt.Printf("It is %d hours %d minutes.", hour, min)
}

//s := 0
//fmt.Scan(&s)
//str := strconv.Itoa(s)
//res :=
