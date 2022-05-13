package main

import "fmt"

/**
任何编程语言都会存在 键值对形式的数据集合  成为字典或映射
最大的特点就是查找速度快，因为其底层是基于哈希表存储的
map是引用类型

 */
func main() {
	//1 初始化
	m := map[string]string{"name":"dhy","age":"24"}
	fmt.Println(m)
	//make函数初始化
	m2 := make(map[int]int,0)
	m2[11] = 12
	fmt.Println(m2)
	//2.常用操作
	//长度
	lens := len(m)
	fmt.Println(lens)
	//添加
	m["temp"] = "temp"
	fmt.Println("添加",m)
	//修改
	m["temp"] = "change"
	fmt.Println("修改",m)
	//删除
	delete(m,"temp")
	fmt.Println("删除",m)
	//查看
	for key,value := range m  {
		fmt.Println(key,value)
	}
}
