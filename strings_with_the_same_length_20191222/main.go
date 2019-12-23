package main

import (
	"fmt"
)

func main() {
	var n int
	//var s, t string
	var s, t, ans string
	fmt.Scan(&n)
	fmt.Scan(&s, &t)
	//var ary []string
	//for i := 0; i < n; i++ {
	//	ary = append(ary, s[i:i+1])
	//  ary = append(ary, t[i:i+1])
	//}
	//res := strings.Join(ary, "")
	//fmt.Println(res)
	for i := 0; i < n; i++ {
		ans += string(s[i]) + string(t[i])
	}
	fmt.Println(ans)
}
