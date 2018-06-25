package main

import "fmt"

//switch 是一个条件语句，用于将表达式的值与可能匹配的选项列表进行比较，并根据匹配情况执行相应的代码块。它可以被认为是替代多个 if else 子句的常用方式。

func muzhis(){
	muzhi := 4
	switch muzhi {
	case 1:
		fmt.Println("shizhi")
	case 2:
		fmt.Println("zhongzhi")
	case 3:
		fmt.Println("wumingzhi")
	case 4:
		fmt.Println("xiaomuzhi")
	}
}

//默认情况
func moren(){
	switch finger := 8; finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("不是指头")
	}
}

//多表达式判断
func duo(){
	letter := "i"
	switch letter {
	case "c":
		fmt.Println("hello")
	case "a","b","d","i": // 一个选项多个表达式
		fmt.Println("world")
	}
}



func main(){
	muzhis()
	moren()
	duo()
}
