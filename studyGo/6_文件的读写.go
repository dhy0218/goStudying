package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func WriteFile(path string){
	//打开文件 新建文件
	f,err := os.Create(path)
	if err!=nil {
		fmt.Println(err)
		return
	}
	//使用完毕 关闭文件
	defer f.Close()
	var s string
	for i:=0; i<20;i++  {
		temp:= strconv.Itoa(i)
		s = "abc" + temp
		f.WriteString(s+"\n")
	}
}
func ReadFile(path string){
	f,err := os.Open(path)
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	buf := make([]byte,100)
	n,err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF{//文件出错同时不是结尾
		fmt.Println(err1)
		return
	}
	fmt.Println(string(buf[:n]))

}
func main() {
	path := "./demo.txt"
	//WriteFile(path)
	ReadFile(path)
}
