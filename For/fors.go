package main

import "fmt"

func fors(){
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d",i)
	}
}

func breaks(){
	for i := 1; i <= 10; i++ {
		//中止条件
		if i > 5 {
			break
		}

		fmt.Printf("%d",i)
	}
}

//continue
//continue 语句用来跳出 for 循环中当前循环。在 continue 语句后的所有的 for 循环语句都不会在本次循环中执行。循环体会在一下次循环中继续执行
func con(){
	for i := 1; i < 10; i++ {
		if i % 2 == 1{
			continue
		}
		fmt.Printf("%d",i)
	}
}

//更多例子
func li1(){
	i := 0
	for ;i <=10; {
		fmt.Printf("%d",i)
		i += 2
	}
}

func li2(){
	i := 0
	for i <= 10 {
		fmt.Printf("%d",i)
		i += 2
	}
}

//无限循环
func forss(){
	for {
		fmt.Println("Hello")
	}
}

func main(){
	fors()
	breaks()
	con()
	li1()
	li2()
	forss()

}
