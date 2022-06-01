package main

import "fmt"

func main() {
    x := MyStruct{"abc"}
    fmt.Println(x.Name)
    x.SetName1("ccc")
    fmt.Println(x.Name)
    x.SetName2("xxx")
    fmt.Println(x.Name)
}

type MyStruct struct {
	Name string
}
//无法修改值
func (s MyStruct) SetName1(name string) {
	s.Name = name
}

func (s *MyStruct) SetName2(name string) {
	s.Name = name
}
