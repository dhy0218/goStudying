package main

import (
	"encoding/json"
	"fmt"
)

type IT struct{
	Company string    `json:"compony"`  //可以改成小写  二次编码
	Subject []string
	IsOk bool
	Price float64
}

func main(){
	s := IT{
		Company: "itcast",
		Subject: []string{"Go","C++","Java"},
		IsOk:    true,
		Price:   666.666,
	}
	//编码，根据内容生成json文本
	//buf,err := json.Marshal(s)
	buf,err := json.MarshalIndent(s,""," ")//格式化编码
	if err!=nil {
		fmt.Println(err)
	}else{
		fmt.Println(string(buf))
	}
}
