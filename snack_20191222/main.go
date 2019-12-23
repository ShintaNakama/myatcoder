package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func max(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// 最大公約数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
  return gcd(b, a%b)
}
// 最小公倍数
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func main() {
	//var a, b, ans int
	var ans int
	//fmt.Scan(&a, &b)
	//d,s := max(a,b)
	sc.Split(bufio.ScanWords)
	s := nextInt()
	d := nextInt()
	//for i := 2; ; i++ {
	//	r := d * i
	//	if r%s == 0 {
	//		ans = r
	//		break
	//	}
	//}
  fmt.Println(gcd(s,d))
	ans = lcm(s, d)

	fmt.Println(ans)
}
