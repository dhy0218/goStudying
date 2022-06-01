package main

import "fmt"

/**
输入：words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
 * 输出：true
 * 解释：在该语言的字母表中，'h' 位于 'l' 之前，所以单词序列是按字典序排列的。
 */
func main() {
	words := []string{"hello","leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words,order))
}
var temp [26]int
func isAlienSorted(words []string, order string) bool {
	for i:=0; i<26;i++  {
		//获取order中字符的位置
		index := order[i] - 'a'
		temp[index] = i
	}
	len := len(words)
	for i:=1;i<len ;i++  {
		//符合字典序应该小于  所以大于的时候就false
		if check(words[i-1],words[i])>0 {
			return false
		}
	}
	return true
}

func check(s1 string, s2 string) int{
	len1 := len(s1)
	len2 := len(s2)
	i := 0
	j := 0
	for {
		if i>=len1||j>=len2 {
			break
		}
		c1 := s1[i] - 'a'
		c2 := s2[j] - 'a'
		if  (c1!=c2){
			return int(temp[c1] - temp[c2])
		}
		i++
		j++
	}
	if i<len1 {
		return 1
	}else if j<len2 {
		return -1
	}
	return 0
}