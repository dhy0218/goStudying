package main

import "fmt"

func main() {
	array := []int{9,5,6,3,4,2,1,7}
	quickSort(&array,0,len(array)-1)
	fmt.Println(array)
}
func quickSort(array *[]int,left int,right int){
	//千万不能忘记返回条件
	if left >= right {
		return
	}
	key := (*array)[left]
	i := left
	j := right
	for ;i<j ;  {
		for ;(i<j)&&(*array)[j]>=key ;  {
			j--
		}
		for ;(i<j)&&(*array)[i]<=key ;  {
			i++
		}
		if i<j {
			temp := (*array)[i]
			(*array)[i] = (*array)[j]
			(*array)[j] = temp
		}
	}
	(*array)[left] = (*array)[i]
	(*array)[i] = key
	quickSort(array,left,i-1)
	quickSort(array,i+1,right)
}
