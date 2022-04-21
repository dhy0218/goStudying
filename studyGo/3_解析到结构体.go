package main

import (
	"encoding/json"
	"fmt"
)

type IT1 struct{
	Company string    `json:"compony"`  //可以改成小写  二次编码
	Subject []string
	IsOk bool
	Price float64
}

func main(){
    jsonBuf := `
{
 "compony": "itcast",
 "Subject": [
  "Go",
  "C++",
  "Java"
 ],
 "IsOk": true,
 "Price": 666.666
}
   `
    var temp IT1
    err := json.Unmarshal([]byte(jsonBuf),&temp)//第二个参数要地址传递
	if err!=nil {
		fmt.Print(err)
	}
    fmt.Println(temp)
}
