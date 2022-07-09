package main

import (
	"fmt"
	"math/rand"
)

type Job struct{
	ID int
	RamdomNum int
}
type Result struct{
	job *Job
	result int
}

func main() {
	jobChan := make(chan *Job,128)
	resultChan := make(chan *Result,128)
	createPool(64, jobChan, resultChan)
	go func(resultChan chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.ID,
				result.job.RamdomNum, result.result)
		}
	}(resultChan)
	//循环创建id 输入到管道
	var id int
	for{
		id++
		r_num := rand.Int()
		job := &Job{
			ID:id,
			RamdomNum:r_num,
		}
		jobChan<-job
	}
}

// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RamdomNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					result: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
