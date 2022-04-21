package main

import "fmt"

type stuxx struct {
	id int
}
func main(){
	i := make([]interface{},3)
	i[0] = "abc"
	i[1] = 1
	i[2] = stuxx{1}

	//for index, data := range i {
	//	  if value,ok := data.(int);ok==true{
	//	  	fmt.Println(value,index,"int")
	//	  }else if value,ok := data.(string);ok==true {
	//		  fmt.Println(value,index,"string")
	//	  }else if  value,ok := data.(stuxx);ok==true{
	//		  fmt.Println(value,index,"stuxx")
	//	  }
	//}

	//switch 类型断言
	for index, data := range i {
		switch value := data.(type) {
		case int: fmt.Println(value,index,"int")
		case string : fmt.Println(value,index,"string")
		case stuxx : fmt.Println(value,index,"stuxx")
		}
	}
}