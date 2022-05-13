package main

import (
	"fmt"
)

func maxRotateFunction(nums []int) int {
	var ans,temp,sum int
	len := len(nums)
	for _,data := range nums  {
		sum += data
	}
	for i:= 0;i<len ;i++  {
		temp += i*nums[i]
	}
	ans = temp
	for i:=1;i<len;i++  {
		x := temp + sum - nums[len-i]*len
		ans = max(x,temp)
		temp = x
	}
	return ans
}

func max(a int, b int) int {
	if a>b {
		return a
	}else{
		return  b
	}
}
func main() {
	var nums = []int{4,3,2,6}
	ans := maxRotateFunction(nums)
	fmt.Println(ans)
}
