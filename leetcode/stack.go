package main

import "fmt"

func main() {
	//x := []int{1,2,3,4,5}
	//x = x[:len(x)-1]
	//fmt.Println(x)
	stack := newStack(10)
	stack.push(1)
	stack.push(2)
	fmt.Println(stack.element)
	fmt.Println(stack.pop())
	fmt.Println(stack.lens())
}

type Stack struct {
	element []int
}

func newStack(val int) *Stack{
	return &Stack{
		element:make([]int,0,val),
	}
}

func (s Stack) lens() int{
	return len(s.element)
}

func (s *Stack)push(val int) {
	s.element = append(s.element,val)
}

func (s *Stack)pop() int{
	if s.lens()==0 {
		return -1
	}
	tmp := s.element[s.lens()-1]
	s.element = s.element[:len(s.element)-1]
	return tmp
}

