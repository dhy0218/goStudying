package main

import "fmt"

func test(){
	func(){
		defer func() {
			if err := recover();err != nil {
				fmt.Println(err)
			}
		}()
		panic("test panic")
		return
	}()
}

//go实现try catch 异常处理

func Try(fun func(),handler func(interface{})){
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func main() {
	//test()
	Try(func() {
		panic("正常执行z中出问题了")
	}, func(i interface{}) {
		fmt.Println(i)
	})
}
