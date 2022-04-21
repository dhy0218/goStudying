package main

import (
	"fmt"
	"regexp"
)

func main(){
	buf := "abc azc aac 223"
	reg1 := regexp.MustCompile("a.c")
	if reg1 == nil {
		fmt.Println("error")
		return
	}
	result1 := reg1.FindAllStringSubmatch(buf,-1)
	fmt.Println(result1)
}
