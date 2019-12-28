package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := nextInt()
	b := nextInt()
	var ans int
	if (b-a)%2 == 0 {
		ans = (b - a) / 2
	} else {
		k := a - 1
		j := n - b
    var min int
		if k < j {
      min = k
		} else {
			min = j
		}
    ans = (min+1) + ((b-a-1)/2)
	}
	fmt.Println(ans)
}
