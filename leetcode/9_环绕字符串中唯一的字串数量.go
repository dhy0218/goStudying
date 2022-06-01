package main

func main() {

}
func findSubstringInWraproundString(p string) (ans int){
	array := []int{}
	k := 0
	//for i := 0;i<len ;i++  {
	//	if i>0 && ((p[i]-p[i-1])+26)%26 == 1 {
	//		k++
	//	}else {
	//		k = 1
	//	}
	//	array[p[i]-'a'] = max(array[p[i]-'a'],k)
	//}
	for index,ch := range p {
		if index>0 && ((byte(ch)-p[index-1])+26)%26==1 {
			k++
		}else {
			k = 1
		}
		array[ch - 'a'] = max(array[ch - 'a'],k)
	}
	for _,data := range array{
		ans += data
	}
	return ans
}