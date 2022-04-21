package main

import (
	"fmt"
	"strings"
)

func main(){
	//Contains 是否包含
	fmt.Println(strings.Contains("hellogo","go"))
	//joins 组合
	s := []string{"abc","go","hello"}
	temp := strings.Join(s,"")
	fmt.Println(temp)
	//index 索引
	fmt.Println(strings.Index("abchello","hello"))
	//repeat
	buf := strings.Repeat("go",3)
	fmt.Println(buf)
	//Split以指定的分隔符拆分
	buf = "hello@go@abc"
	s2 := strings.Split(buf,"@")
	fmt.Println(s2)
	fmt.Println(len(s2))
	//trim去掉两头的空格
	b := strings.Trim("  are  you ok "," ")
	fmt.Println(b)
	//feilds 去掉所有空格  生成一个切片
	c := strings.Fields("  are u ok ")
	fmt.Println(c)
	for index,data := range c{
		fmt.Println(index,data)
	}
}
