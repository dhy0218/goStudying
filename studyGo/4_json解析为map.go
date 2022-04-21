package main

import (
	"encoding/json"
	"fmt"
)

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
	//创建map
	m := make(map[string]interface{},4)
	err := json.Unmarshal([]byte(jsonBuf),&m)
	if err!=nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v",m)
}
