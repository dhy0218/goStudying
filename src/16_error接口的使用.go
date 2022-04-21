package main

import (
	"errors"
	"fmt"
)

func main(){
	errr := fmt.Errorf("%s","this is a normal error")
	fmt.Println(errr)

	err2 := errors.New("this is a err2")
	fmt.Println(err2)
}
