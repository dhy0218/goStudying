package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var randNum int
	//创造随机数 四位的
	CreatNum(&randNum)
	//保存每一位数字
	silceMy := make([]int,4)
	getNum(silceMy,randNum)
	fmt.Println(randNum)
	fmt.Println(silceMy)
	fmt.Println("**************")
}

func getNum(my []int, num int) {
	my[0] = num/1000
	my[1] = num%1000 / 100
	my[2] = num %100 / 10
	my[3] = num %10
}

func CreatNum(i *int) {
	rand.Seed(time.Now().UnixNano())
	var num int
	for{
		num = rand.Intn(10000)
		if num>=1000 {
			break
		}
	}
	*i = num
	fmt.Println(num)
}
