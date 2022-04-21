package main

import "fmt"

//切片与底层数组的关系 另外的新切片 对切片的操作回影响原数组
//切片的截取
func main(){
	array := []int{0,1,2,3,4,5,6,7,8,9}
	s1 := array[:]
	fmt.Println(s1,len(s1),cap(s1))
	//操作某个元素 和数组操作方式一样
	data := array[0]
	fmt.Println(data)

	s2 := array[3:6:7]
	fmt.Println(s2)

	s3 := array[:6]//从0开始取六个元素 容量也是6
	fmt.Println(s3)
	fmt.Println("****************************")
	Mytest()
	fmt.Println("**********扩容特点 两倍")
	test := array[0:3:3]
	oldCap := cap(test)
	fmt.Println(oldCap)
	test = append(test,1 )
	newCap := cap(test)
	fmt.Println(newCap)

}
//测试常用函数 append copy
func Mytest(){
	s1 := []int{}
	s1 = append(s1, 1)
	fmt.Println(s1, len(s1),cap(s1))
    //如果超过原来容量 通常以2倍容量扩容
}

func SliceTest(){
	a := []int{1,2,3,4,5}
	//low high max
	s := a[0:3:5]
	fmt.Println(s)
	fmt.Println(len(s))
    fmt.Println(cap(s))
	//切片和数组的区别
	//数组【】里面的长度是固定的一个常量，数组不能修改长度 len和cap是固定的
	b := [5]int{}
	//切片【】里面为空或者为。。。 切片的长度可以不固定
	x := []int{}
	fmt.Println(b,x)
}
//切片的创建
func Test2(){
	//自动推导类型同时初始化
	s1 := []int{1,2,3,4}
	fmt.Println(s1)

	//借助make函数 格式 make{切片类型，长度，容量}
	s2 := make([]int,5,10)
	fmt.Println(s2)

	//借助make函数 格式 make{切片类型，长度，容量}  容量不固定可以不写容量
	s3 := make([]int,5)
	fmt.Println(s3)
}
