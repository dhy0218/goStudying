package main

import (
	"fmt"
	"time"
)

func newTask(){
	for {
		fmt.Println("this is a new task")
		time.Sleep(time.Second)
	}
}
func main() {

	go newTask()//新建一个协程
	for   {
		fmt.Println("this is a main task")
		time.Sleep(time.Second)

	}
}
