package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 金钱卦起卦法，每次随机出现三个数字（2或3），取三数之和。
func main() {
	rand.Seed(time.Now().UnixNano())
	//	for i := 0; i < 3; i++ { //产生三个数字
	//		fmt.Println(rand.Intn(2)) //数字是0或1
	//	}
	for i := 0; i < 6; i++ { //产生六组数字,每组三个
		var num1 = rand.Intn(2) + 2
		var num2 = rand.Intn(2) + 2
		var num3 = rand.Intn(2) + 2
		//		defer fmt.Println(num1 + num2 + num3)
		//		fmt.Println(num1 + num2 + num3) //按顺序检查数字

		//逆序依次产生六爻
		var sum = num1 + num2 + num3
		if sum%2 == 0 {
			defer fmt.Println("- -", sum) //阴爻
		} else {
			//		if num%2 != 0 {
			defer fmt.Println("___", sum) //阳爻
		}
	}
}
