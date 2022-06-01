package main

func main() {
	array := []int{1,2,3,4,1,1,1,1}
	repeatedNTimes(array)
}
func repeatedNTimes(nums []int) int{
	m := map[int]bool{}
	for _,data := range nums  {
		if m[data] {
			return data
		}
		m[data] = true
	}
	return -1
}
