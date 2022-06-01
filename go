1.判断类型用reflect.TypeOf() str[0] 类型为uint8  range时类型为int32  byte(range)类型为uint8
2.go数据类型分为四类：basic types基础数据类型  aggregatetypes 复合数据类型  reference types引用数据类型 interfacetypes 接口数据类型
 基础：数字 字符串 布尔
 复合数据类型：数组、结构体
 引用：指针 slice map channel 接口和函数
 3.java byte 1
        short 2
        int 4
        long 8
        float 4
        double 8
        boolean 1
        char 2
    go  int8 1
        int16 2
        int32 4
        int64 8
        float32 4
        float64 8

 4.goBasic中的结构体方法和结构体指针表明 当需要修改结构体的值时，需要使用结构体指针
