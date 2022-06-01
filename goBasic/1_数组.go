package main

import "fmt"

func test1(array [3]int){
	array[1] = 3
}
func test2(array *[3]int){
	(*array)[1] = 3
}

func main() {
	//1.初始化方式
	var array1 [3]int = [3]int{1,2,3}
	fmt.Println(array1)

	var array2 = [3]int{4,5,6}
	fmt.Println(array2)

	var array3 = [...]int{7,8,9}
	fmt.Println(array3)
    //2.遍历方式
	for index,data := range array3  {
		fmt.Println(index,data)
	}
	//3.int 初始0  string “”  bool false

    //go函数数组的值传递与引用传递
    array := [3]int{1,2,3}
    fmt.Println(array)
    test1(array)
    fmt.Println(array)
    //可以通过指针形式 传入参数 修改数组值
    test2(&array)
    fmt.Println(array)//1,3,3

    fmt.Println("测试语法糖【：i】********************")
    testArrays := [][]int{
    	{1,2},
    	{3,4},
        {5,6},
        {7,8},
    }
	for i,data1 := range testArrays  {
		for j,data2 := range testArrays[:i]{
			for k,data3 := range testArrays[:j]  {
				fmt.Println(i,data1,j,data2,k,data3)
			}
		}
	}

}
