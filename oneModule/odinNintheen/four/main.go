package main

import "fmt"

func main() {
	s, ch, res := 0, 0, 0
	fmt.Scan(&s)
	ch = s / 1000
	pr, tw, thr := ch/100, ch/10, ch/1
	tw = tw % 10
	thr = thr % 10
	//fmt.Println(pr, tw, thr)
	res = pr + tw + thr
	//fmt.Println(res)
	ch = s % 1000

	pr, tw, thr = ch/100, ch/10, ch/1
	tw = tw % 10
	thr = thr % 10
	res2 := pr + tw + thr
	//fmt.Println(res2)
	if res == res2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
