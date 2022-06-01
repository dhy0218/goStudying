package main

import (
	"fmt"
	"sort"
)

/**
输入：intervals = [[3,4],[2,3],[1,2]]
输出：[-1,0,1]
 */
func main() {
	intervals := [][]int{{3,4},{2,3},{1,2}}
	fmt.Println(findRightInterval(intervals))
}
func findRightInterval(intervals [][]int) []int {
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
		fmt.Println(intervals[i])
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	fmt.Println(intervals)
	n := len(intervals)
	ans := make([]int, n)
	for _, p := range intervals {
		i := sort.Search(n, func(i int) bool { return intervals[i][0] >= p[1] })
		if i < n {
			ans[p[2]] = intervals[i][2]
		} else {
			ans[p[2]] = -1
		}
	}
	return ans

}
