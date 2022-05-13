package main

import (
	"fmt"
	"sort"
)

/**
切片 动态数组
每个切片对象都维护着数组指针、切片长度、切片容量三个数据
1.1 创建切片
var num []int
var data []int{11,22}
data := []int{11,22}

make只用于 切片 字典 channel  可以设置数据个数和容量
1.2 理解数组指针
1.3自动扩容
   当向切片中追加的数据个数大于容量时，内部回自动扩容每次都会新建数组，扩容时当前容量的2倍
1.4 copy函数
1.5  sort.slice（）排序

***************************************************************************
切片的坑
切片作为函数的形参是 引用传递  传递的是指向切片的内存地址  这意味着函数中对切片的任何操作都会在函数结束后体现出来，另外，函数中传递切片要比
传递数组更加高效    但是数组会在go中传递数值 也就是说会拷贝整个数组的值

使用make 创建切片时，len设置为0  否则在0~len之间的数据会有默认值出现
 */

type stu struct{
	person string
	height int
	weight int
}
func main() {
	//make 类型 长度 容量
	var v1 = make([]int,1,2)
	fmt.Println(v1) //0
	//append可以在切片后追加
	v2 := append(v1, 12)
	fmt.Println(v2)// 0，12

	v1[0] = 33
	fmt.Println(v1)//33
	fmt.Println(v2)//33 12
    //copy  要复制的目标slice  源slice
    b1 := []int{1,2,3,4,5,6}
    b2 := []int{7,8,9}
    copy(b1,b2)
    fmt.Println(b1) //7 8 9 4 5 6
    fmt.Println(b2) // 7 8 9
    //测试sort
    students := make([]stu,0)
    students = append(students,stu{"abc",11,33},stu{"zxc",22,55},stu{"uio",111,45})
    fmt.Println(students)
    //身高由小到大顺利排列
    sort.Slice(students, func(i, j int) bool {
		return students[i].height < students[j].height
	})
    fmt.Println(students)
}
