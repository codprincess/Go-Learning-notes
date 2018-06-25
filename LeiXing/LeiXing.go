package main

import "fmt"

/*笔记:下面是 Go 支持的基本类型：

bool类型表示一个布尔值即真或者假，值为true或者false

数字类型
int8, int16, int32, int64, int
uint8, uint16, uint32, uint64, uint
float32, float64
complex64, complex128
byte
rune

string*/

func bools(){
	a := true
	b := false

	fmt.Println("a:",a,"b",b)
	c := a && b //c 赋值为 a && b。仅当 a 和 b 都为 true 时，操作符 && 才返回 true。因此，在这里 c 为 false
	fmt.Println("c:",c)
	d := a || b //当 a 或者 b 为 true 时，操作符 || 返回 true。在这里，由于 a 为 true，因此 d 也为 true
	fmt.Println("d:",d)
}

//类型转换
func lei(){
	i := 55 //int
	j := 66.6 //float64
	sum := i + int(j) //不允许int + float64
	fmt.Println(sum)
}


func main(){
	bools()
	lei()
}
