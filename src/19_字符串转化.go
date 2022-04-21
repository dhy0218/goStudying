package main

import (
	"fmt"
	"strconv"
)

func main(){
	//转换为字符串后追加到字节数组
	slice := make([]byte,0,1024)
	slice = strconv.AppendBool(slice,true)
	slice = strconv.AppendInt(slice,1033,10)
	slice = strconv.AppendQuote(slice,"abchello")
	fmt.Println(string(slice))//需要转换为string后在打印
	//其他类型转换为字符串
	str := strconv.FormatBool(true)
	fmt.Println(str)
	//整型转字符串
	str = strconv.Itoa(66)
	fmt.Println(str)
	//字符串转其他类型
	var flag bool
	var er error
	flag,er = strconv.ParseBool("true")
	if er == nil {
		fmt.Println(flag)
	}else {
		fmt.Println(er)
	}
}
