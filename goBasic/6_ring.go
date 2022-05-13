package main

import (
	"container/ring"
	"fmt"
)
/**
Ring类型的数据结构仅由它自身即可代表，而List类型则需要由它以及Element类型联合表示。这是表示方式上的不同，也是结构复杂度上的不同。
一个Ring类型的值严格来讲，只代表了其所属的循环链表中的一个元素，而一个List类型的值则代表了一个完整的链表。这是表示维度上的不同。
在创建并初始化一个Ring值的时候，我们可以指定它包含的元素的数量，但是对于一个List值来说却不能这样做（也没有必要这样做）。循环链表一旦被创建，其长度是不可变的。这是两个代码包中的New函数在功能上的不同，也是两个类型在初始化值方面的第一个不同。
仅通过var r ring.Ring语句声明的r将会是一个长度为1的循环链表，而List类型的零值则是一个长度为0的链表。别忘了List中的根元素不会持有实际元素值，因此计算长度时不会包含它。这是两个类型在初始化值方面的第二个不同。
Ring值的Len方法的算法复杂度是 O(N) 的，而List值的Len方法的算法复杂度则是 O(1) 的。这是两者在性能方面最显而易见的差别。
 */
func main() {
	size := 10
	myRing := ring.New(size)
	fmt.Println("empty ring",myRing)
	for i := 0;i<size ; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}
	printRing(myRing)
	myRing = myRing.Move(-1)
	printRing(myRing)
}

func printRing(myRing *ring.Ring) {
	fmt.Println(myRing.Len())
	len := myRing.Len()
	for i:=0;i<len ;i++  {
		fmt.Println(myRing.Value)
		myRing = myRing.Next()
	}

}