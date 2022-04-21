package main

import "fmt"

type Father interface {
	sayHi()
}
type child interface {
	Father//匿名字段继承了sayHi
	sing(x string)
}
type stud struct{
	name string
	id int
}
func (temp *stud)sayHi(){
	fmt.Println(temp.name,temp.id)
}
func (temp *stud)sing(i string){
	fmt.Println(i)
}
func main(){
	var c child
	s := &stud{"name",124}
	c = s
	c.sayHi()
	c.sing("aaaa")
}
