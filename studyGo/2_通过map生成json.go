package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	m := make(map[string]interface{},4)
	m["company"] = "itcast"
	m["subjects"] = []string{"Go","Java"}
	m["isOk"] = true
	m["price"] = 666.666

	//编码成json
	result,err := json.Marshal(m)
	if err!= nil {
		fmt.Println(err)
	}else{
		fmt.Println(string(result))
	}
}
