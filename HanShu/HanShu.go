package main

import "fmt"

/*func functionname(parametername type) returntype {
// 函数体（具体实现的功能）
}

函数的声明以关键词 func 开始，后面紧跟自定义的函数名 functionname (函数名)。函数的参数列表定义在 ( 和 ) 之间，返回值的类型则定义在之后的 returntype (返回值类型)处。声明一个参数的语法采用 参数名 参数类型 的方式，任意多个参数采用类似 (parameter1 type, parameter2 type) 即(参数1 参数1的类型,参数2 参数2的类型)的形式指定。之后包含在 { 和 } 之间的代码，就是函数体。*/

//计算商品价格函数
func calculateBill(price int, no int) int {
	var totalPrice = price * no //商品总价 = 商品单价 * 数量
	return totalPrice //返回总价 类型为int
}

//Go语言支持一个函数有多个返回值
func rectprops(length,width float64)(float64,float64){
	var area = length * width
	var perimeter = (length + width) * 2
	return area,perimeter
}

//命名返回值 一旦命名了返回值，可以认为这些值在函数第一行就被声明为变量了
func rect(length,width float64)(area, perimeter float64){
	area = length * width
	perimeter = (length + width) * 2
	return
}

//空白符
// _ 在 Go 中被用作空白符，可以用作表示任何类型的任何值。
/*以 re 函数为例，该函数计算的是面积和周长。假使我们只需要计算面积，而并不关心周长的计算结果，该怎么调用这个函数呢？这时，空白符 _ 就上场了*/
func re(length,width float64)(area,per float64){
	area = length * width
	per = (length + width) * 2
	return
}

func main(){
	/*a := calculateBill(5,10)
	fmt.Println(a)*/

	/*area,perimeter := rectprops(10.8,5.6)
	fmt.Println(area,perimeter)*/

	/*area2,perimeter2 := rect(10.8,5.6)
	fmt.Println(area2,perimeter2)*/

	area3,_ := re(10.8,5.6)
	fmt.Println(area3)
}
