package rectangle

//自定义包

import (
	"math"
	"fmt"
)

/*我们将组织代码，使得所有与矩形有关的功能都放入 rectangle 包中。

我们会创建一个自定义包 rectangle，它有一个计算矩形的面积和对角线的函数。

属于某一个包的源文件都应该放置于一个单独命名的文件夹里。按照 Go 的惯例，应该用包名命名该文件夹。

因此，我们在 geometry 文件夹中，创建一个命名为 rectangle 的文件夹。在 rectangle 文件夹中，所有文件都会以 package rectangle 作为开头，因为它们都属于 rectangle 包*/

/*将 rectangle 包中的函数 Area 和 Diagonal 首字母大写。在 Go 中这具有特殊意义。在 Go 中，任何以大写字母开头的变量或者函数都是被导出的名字。其它包只能访问被导出的函数和变量。在这里，我们需要在 main 包中访问 Area 和 Diagonal 函数，因此会将它们的首字母大写。*/

//init函数 init 函数可用于执行初始化任务，也可用于在开始执行之前验证程序的正确性
func init(){
	fmt.Println("hello")
}


func Area(len,wid float64) (area float64){
	area = len * wid
	return
}

func Diagonal(len,wid float64)(diagonal float64){
	diagonal = math.Sqrt((len * len) + (wid * wid)) //math 包下面的 Sqrt 函数用于计算平方根
	return
}