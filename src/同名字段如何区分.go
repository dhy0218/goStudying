package main
import "fmt"

type Person1 struct{
	name string
	sex byte
	age int
}

type Student2 struct{
	Person1//只有类型 没有名字 匿名字段 继承了person的成员
	id int
	addr string
	name string
}

func main(){
	var s Student2
	//默认规则 就近原则  如果能在本作用域找到此成员就操作此成员
	    //                没有找到 就找到继承字段
	s.name = "sss"
	fmt.Println(s)
}
