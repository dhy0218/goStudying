package main

import "fmt"

type Person11 struct {
	name string
}

func (p Person11) SetPerson(){
	fmt.Println("setPerson")
}

func (p *Person11) SetPersonInfo(){
	fmt.Println("setPersonInfo")
}

func main(){
	//结构体变量是一个指针，它能够调用哪些方法，这些方法的集合成为方法集
	p := &Person11{"aaaa"}
	p.SetPerson()//内部做的转换  （*p）
	p.SetPersonInfo()
	(*p).SetPersonInfo()
	(*p).SetPerson()
}