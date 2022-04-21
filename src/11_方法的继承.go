package main

import "fmt"

type Persontest struct {
	name string
}
type Strudenttest struct{
	Persontest
	addr string
}
func (temp Persontest)printinfo(){
	fmt.Println(temp.name)
}
//student也实现了print方法 叫重写 就近原则
func (temp *Strudenttest) printinfo(){
	fmt.Println(temp)
}
func main(){
	s := &Strudenttest{
		Persontest: Persontest{name:"abc"},
		addr:       "dddd",
	}
	//方法的继承
	s.printinfo()
}
