package testPackage

import "strings"

func SplitHelp(s,seb string)(result []string){
	i := strings.Index(s,seb)
	for ;i>-1 ;  {
		result = append(result,s[:i])
		s = s[i+(len(seb)):]
		i = strings.Index(s,seb)
	}
	//最后一个就没有":"
	result = append(result,s)
	return result
}


