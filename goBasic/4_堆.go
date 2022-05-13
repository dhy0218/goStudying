package main

import (
	"container/heap"
	"fmt"
)

/**
int  32位系统是4字节 64位系统是8字节
int8 一字节 -128~127
int16 2字节
int32 4字节
int64 8字节
 */
type student struct{
	name string
	age int32
	id int64
}

type heapStudents []student

func (h *heapStudents) Len() int {
	return len(*h)
}

func (h *heapStudents) Less(i, j int) bool {
	return (*h)[i].age > (*h)[j].age //大顶堆
}

func (h heapStudents) Swap(i, j int) {
	h[i],h[j] = h[j],h[i]
}

func (h *heapStudents) Pop() interface{} {
	old := *h
	val := old[0]
	*h = old[1:] // 每次pop都会新开辟一个切片, 如果需要这里可以做一个优化
	heap.Init(h)
	return val
	panic("implement me")
}

func (h *heapStudents) Push(x interface{}){
	val,err := x.(student)
	if err != true {
		fmt.Println(fmt.Errorf("value error"))
	}
	*h = append(*h,val)//每次push都会新开辟一个切片
	heap.Init(h)
}
func main() {
	myHeap := heapStudents{}
	myHeap = append(myHeap,
		student{"a",1,2},
		student{"b",2,3},
		student{"c",4,5},
		student{"d",6,7},
	)
	fmt.Println(myHeap)

	fmt.Println("heap init")
	myHeapInit := &myHeap
	heap.Init(myHeapInit)
	fmt.Println(myHeap)

	fmt.Println("push")
	myHeap.Push(student{"h",4,5})
	fmt.Println(myHeap)

	fmt.Println("pop")
	fmt.Println(len(myHeap))
	len2 := len(myHeap)
	for i:=0;i<len2;i++  {
		fmt.Println(myHeap.Pop())
	}

}
