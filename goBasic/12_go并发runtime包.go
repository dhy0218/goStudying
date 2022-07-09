package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
1.gosched 让出时间片
2.goexit 退出当前线程
3.GOMAXPROCS 设置核心数实现并行效果
 */
func main() {
	/**
	go func(s string) {
		for i:=0;i<2 ;i++  {
			fmt.Println(s)
		}
	}("world")
	for i:=0;i<2;i++  {
		//让出cpu时间片
		runtime.Gosched()
		fmt.Println("hello")
	}
	 */
	/** print B.defer A.defer 直接退出后面的都不运行了
	go func() {
		defer fmt.Println("A.defer")
		func(){
			defer fmt.Println("B.defer")
			//结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for  {

	}
	 */
	runtime.GOMAXPROCS(8)
	go a()
	go b()
	time.Sleep(time.Second)

}
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}