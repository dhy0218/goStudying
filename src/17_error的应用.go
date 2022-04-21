package main

import (
	"errors"
	"fmt"
)

func Mydiv(a,b int)(result int,err error){
	if b==0 {
		err = errors.New("分母不能为0")
	}else {
		result =  a/b
	}
	return
}
func main(){
	result,err := Mydiv(5,10)
	if err!=nil {
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
}
