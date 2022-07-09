package main

import "fmt"

/**
1.go语言的并发模型是csp 提倡通过通信共享内存而不是通过共享内存来是实现通信
如果说goroutine是go程序并发的执行体，channel就是他们之间的连接
2channel可以让一个goroutine发送特定值到另一个goroutine的通信机制
3channel操作有发送接收和关闭三个操作
4如果设置的是一个无缓冲的通道 发送接收就会deadlock
解决的办法 一是开启一个协程接收 二是有缓冲的通道
 */
var ch1 chan int //声明一个传递整型的通道
var ch2 chan bool
var ch3 chan []int

func main() {
	/**
	无缓冲的通道会阻塞
	ch := make(chan int)
	//发送
	ch <- 10//把10发送到ch
	//接收
	x := <- ch//从ch中接收值并赋值给变量x
	fmt.Println(x)
	close(ch)
	 */
    //一个解决方案就是开启一个协程 执行接收操作
    //必须要先开启一个go协程 来接收通道的东西
    ch := make(chan int)
    go receive(ch)
	ch<- 123
	fmt.Println("发送成功")
	fmt.Println("+++++++++++++++++++++++++++++++++++++")
	//有缓冲
	chx := make(chan int,1)
	chx <- 456
	fmt.Println("发送成功")
	temp := <- chx
	fmt.Println(temp)
	fmt.Println("接收成功")
	//无缓冲通道
	testchan := make(chan int)
	go func() {
		for i:=0;i<5 ;i++  {
			testchan<- i
		}
		close(testchan)
	}()
	for{
		if data,ok := <-testchan;ok {
			fmt.Println(data)
		}else{
			break
		}
	}
	fmt.Println("main over")
}
func receive(rec chan int) {
	temp := <- rec
	fmt.Println("接收成功",temp)
}