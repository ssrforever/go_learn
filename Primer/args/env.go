// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	Echo1()
	Echo2()
}

func Echo1() {
	start := time.Now() // 记录开始时间
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Since(start) // 计算耗时
	fmt.Println("耗时：", elapsed)
}
func Echo2() {
	start := time.Now() // 记录开始时间
	fmt.Println(strings.Join(os.Args[1:], " "))
	elapsed := time.Since(start) // 计算耗时
	fmt.Println("耗时：", elapsed)
}
