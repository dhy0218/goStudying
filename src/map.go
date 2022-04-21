package main

import "fmt"

func main(){
	//map是无序的
	var m1 map[int]string
	fmt.Println(m1)
	//可以通过make 创建
	m2 := make(map[int]string)
	fmt.Println(m2)
	//指定长度
	m3 := make(map[int]string,10)
	fmt.Println(m3)
	m3[1] = "mike"
	m3[2] = "go"
	fmt.Println(m3)
	//赋值修改内容
	fmt.Println("修改后的*************")
	m3[1] = "test"
	fmt.Println(m3)
	//map的遍历
	for key,value := range m3   {
		fmt.Println(key,value)
	}
	//判断key是否存在
	value,ok := m3[0]
	if ok == true {
		fmt.Println(value)
	}else {
		fmt.Println("bucunzai")
	}
	//删除
	fmt.Println("删除key")
	delete(m3,1)
	fmt.Println(m3)
}
