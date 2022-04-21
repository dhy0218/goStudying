package main

import "fmt"

type Person struct{
	name string
	sex byte
	age int
}

type Student1 struct{
	Person//只有类型 没有名字 匿名字段 继承了person的成员
	id int
	addr string
}

func main(){
	//成员初始化
	//顺序初始化
	var s1 Student1 = Student1{Person{"mike",'m',18},1,"aaaa"}
	fmt.Println(s1)
	//自动推导类型
	s2 := Student1{Person{"mike",'m',18},1,"aaaa"}
	fmt.Println(s2)
	fmt.Printf("%+v",s2)
	//指定成员初始化
	s3 := Student1{id:1}
	fmt.Println(s3)
	//。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。
    s3.id = 2
}
