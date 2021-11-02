package main

import (
	"fmt"
	"math"
)

//常量

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 50000

	const d = 3e20 / n
	fmt.Println("值为：", d, "类型为：", fmt.Sprintf("%T", d))
	//转成int64类型
	fmt.Println("值为：", int64(d), "d的类型是：", fmt.Sprintf("%T", int64(d)))

	//math.Sin 返回正弦值
	fmt.Println(math.Sin(n))
}
