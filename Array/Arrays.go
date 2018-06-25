package main

import "fmt"

//数组是同一类型元素的集合。例如，整数集合 5,8,9,79,76 形成一个数组。Go 语言中不允许混合不同类型的元素，例如包含字符串和整数的数组。注：如果是 interface{} 类型数组，可以包含任意类型

//数组的声明
//一个数组的表示形式为 [n]T。n 表示数组中元素的数量，T 代表每个元素的类型。元素的数量 n 也是该类型的一部分（稍后我们将详细讨论这一点）

func li1(){
	var a [3]int
	a[0] = 12
	a[1] = 78
	a[2] = 50
	fmt.Println(a)
}

func li2(){
	//var a = [3]int{12,78,50}
	a := [3]int{12,78,50}
	fmt.Println(a)
}

//在简略申明中，不需要将数组中的所有元素赋值
func li3(){
	a := [3]int{12}
	fmt.Println(a)
}

//忽略数组长度，用 ... 代替，让编译器自动为你计算长度
func li4(){
	a := [...]int{12,78,50,100}
	fmt.Println(a)
}

//数组的大小是类型的一部分，所以[5]int和[25]int是不同的类型
//数组不能调整大小，不要担心这个限制，因为 slices 的存在能解决这个问题

//数组是值类型
//Go 中的数组是值类型而不是引用类型。这意味着当数组赋值给一个新的变量时，该变量会得到一个原始数组的一个副本。如果对新变量进行更改，则不会影响原始数组

func li5(){
	a := [...]string{"a","b","c","d","e","f","g"}
	b := a
	b[0] = "z"
	fmt.Println("a is",a)
	fmt.Println("b is",b)
}

//同样，当数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变。
var arrs = [5]int{1,2,3,4,5}

func li6(arrs [5]int){
	arrs[0] = 55
	fmt.Println("函数里的num是",arrs)
}

//数组的长度
//通过将数组作为参数传递给 len 函数，可以得到数组的长度
func li7(){
	a := [...]int{1,2,3,4,5,6}
	fmt.Println("数组的长度",len(a))
}

//使用range迭代数组
//for循环遍历数组中的元素
func li8(){
	a := [...]float64{67.7,89.8,21,78,66}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])//注意占位符
	}
}
//for i, v := range a 利用的是 for 循环 range 方式。 它将返回索引和该索引处的值。 我们打印这些值，并计算数组 a 中所有元素的总和
//如果你只需要值并希望忽略索引，则可以通过用 _ 空白标识符替换索引来执行
/*for _, v := range a { // ignores index
}*/
func li9(){
	a := [...]float64{67.7,89.8,21,78,66}
	sum := float64(0)

	for i,v := range a {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a",sum)
}

//多维数组
func li10(a [3][2]string){
	for _,v1 := range a {
		for _,v2 := range v1 {
			fmt.Printf("%s",v2)
		}
		fmt.Printf("\n")
	}
}
func li11(){
	a := [3][2]string{
		{"lion","tiger"},
		{"cat","dog"},
		{"pigeon","peacock"},
	}
	li10(a)
	fmt.Printf("\n")
	var b = [3][2]string{
		{"a","zz"},
		{"c","c2"},
		{"d1","d2"},
	}
	li10(b)

}

func main(){
	li1()
	li2()
	li3()
	li4()
	li5()
	li6(arrs)
	fmt.Println(arrs)
	li7()
	li8()
	li9()
	li11()
}
