package main

func main(){

	//第一种写法
	/*num := 10
	if num % 2 == 0 {
		fmt.Println("这是偶数")
	}else{
		fmt.Println("这不是偶数")
	}*/

	//第二种写法
	//num 在 if 语句中进行初始化，num 只能从 if 和 else 中访问。也就是说 num 的范围仅限于 if else 代码块。如果我们试图从其他外部的 if 或者 else 访问 num,编译器会不通过
	/*if num := 10;num % 2 == 0 {
		fmt.Println(num,",是偶数")
	}else{
		fmt.Println(num,",不是偶数")
	}*/
}
