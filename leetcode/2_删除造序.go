package main

import (
	"fmt"
	"reflect"
	_"reflect"
)

func main() {
	str := "abc"
	fmt.Println(reflect.TypeOf(str[0]))
	for index,temp := range str  {
		fmt.Println(index,reflect.TypeOf(byte(temp)))
	}
	//strs := [3]string{"abf","bce","cae"}
	//ans := 0
	//fmt.Println(strs[0])
	//for j := range strs[0] {
	//	for i := 1; i < len(strs); i++ {
	//		if strs[i-1][j] > strs[i][j] {
	//			ans++
	//			break
	//		}
	//	}
	//}

}
