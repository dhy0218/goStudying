package main

import (
	"fmt"
	"time"
)
func printer(s string){
	for _,data := range s{
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
}
func person1(){
	printer("hello")
}
func person2(){
	printer("world")
}
func main(){

	go person1()
	go person2()
	for {

	}

}
