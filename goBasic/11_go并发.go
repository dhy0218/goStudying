package main

/**
goroutine奉行通过通信来共享内存，而不是共享内存来通信
 */
import (
	"fmt"
	"sync"
	"time"
)

func hello(){
	fmt.Println("hello")
}
//启动多个goroutine
var wg sync.WaitGroup

func hello1(i int){
	defer wg.Done()
	fmt.Println("hello Goroutine",i)
}
//func main() {
//	//go hello()
//	//fmt.Println("main")
//	for i:=0;i<10;i++{
//		wg.Add(1)
//		go hello1(i)
//	}
//	wg.Wait()
//}
//主协程退出 其他任务就会自动退出
func main() {
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}