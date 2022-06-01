package main

import "fmt"

/**
输入：digits = "23"
 * 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
 */
func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
var m map[string]string
var ans []string
var temp6 string
func letterCombinations(digits string) []string {
	m = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	if digits == "" {
		return ans
	}
	dfs(digits,0)
	return  ans
}

func dfs(digits string, i int) {
	if len(digits)== len(temp6) {
		ans = append(ans, temp6)
		return
	}
	indexString := string(digits[i])
	letters := m[indexString]
	lens := len(letters)
	for j:=0;j<lens ;j++  {
		temp6 = temp6 + string(letters[j])
		dfs(digits,i+1)
		help(&temp6)
	}
}

func help(s *string) {
	lens := len(*s)
	var ss string = ""
	for i := 0;i<lens-1 ;i++  {
		ss = ss + string((*s)[i])
	}
	*s = ss
}