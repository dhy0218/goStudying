package main

import "fmt"

type Humaner interface {
	sayHi()
}
type stu struct {
	name string
	id int
}
func (temp *stu)sayHi(){
	fmt.Println(temp.name,temp.id)
}
type teacher struct {
	add string
}
func (temp *teacher)sayHi(){
	fmt.Println(temp.add)
}

//定义一个普通函数 函数的参数为接口类型
func whoSayHi(h Humaner){
	h.sayHi()
}


func main(){
	//定义接口类型的变量
	var i Humaner
	//只要实现了该方法的类型  都可以对i赋值
	s := &stu{"abc",111}
	i = s
	t := &teacher{add:"addddd"}
	i = t
	i.sayHi()
	//多态
	whoSayHi(s)
	whoSayHi(t)
}
func main1(){
	//定义接口类型的变量
	var i Humaner
	//只要实现了该方法的类型  都可以对i赋值
	s := &stu{"abc",111}
	i = s
	i.sayHi()
	whoSayHi(s)

	t := &teacher{add:"addddd"}
	i = t
	i.sayHi()
	whoSayHi(t)
}
