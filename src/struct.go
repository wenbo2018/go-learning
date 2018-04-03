package main

import (
	"fmt"
)

//定义一个person结构，类似于在PHP定义了一个person的class
type person struct {
	name string //默认零值为空字符串
	age  int    //默认零值为0
}

func main() {
	//如何传递指针
	a := &person{ //变量初始化时，加&取地址
		name: "GO",
		age:  8} //注意，如果没有逗号，则}不能另起新行，否则会报错
	//当传地址的时候，想操作属性时，Go语言可以不需要加*，直接操作
	a.age = 9
	fmt.Println("我在A前面输出", a)
	A(a)
	fmt.Println("我在A后面输出", a) //很显然,A里面的赋值并没有改变a的值，证明结构是值类型，传值是值拷贝
}

//结构也是值类型，传值的时候也是值拷贝
func A(per *person) { //此时为指针类型
	per.name = "PHP"
	fmt.Println("我在A里面输出", per)
}
