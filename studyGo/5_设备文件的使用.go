package main

import (
	"fmt"
	"os"
)

func main(){
	//os.Stdout.Close()  //关闭后无法输出
	fmt.Println("are u ok?")//往标准输出设备屏幕写内容
	//标准设备文件 os.Stdout 默认已经打开，用户可以直接使用
	os.Stdout.WriteString("abc")

	var a int
	fmt.Scan(&a)
	fmt.Println(a)
}
