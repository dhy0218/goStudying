package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	//将0到100发送到ch1
	go func() {
		for i:=0;i<100 ;i++  {
			ch1 <- i
		}
		close(ch1)
	}()
	//从ch1接收值 并将该值的平方发送到ch2
	go func() {
		for{
			i,ok := <-ch1
			if !ok {
				break
			}
			ch2<-i*i
		}
		close(ch2)
	}()
	for i := range ch2{
		fmt.Println(i)
	}
}
