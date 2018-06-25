package main

import (
	"./rectangle" //我们必须指定自定义包相对于工作区内 src 文件夹的相对路径  导入自定义包
	"fmt"
	"log"
)
//使用空白符
//导入了包，却不在代码中使用它，这在 Go 中是非法的(原因是为了避免导入过多未使用的包，从而导致编译时间显著增加)
//然而，在程序开发的活跃阶段，又常常会先导入包，而暂不使用它。遇到这种情况就可以使用空白标识符 _
/*package main
import (
_ "geometry/rectangle"
)
func main() {
}*/



var rectlen,rectwid float64 = 6,7

//验证长度和宽度不能小于0
func init(){
	println("nihao")
	if rectlen < 0 {
		log.Fatal("长度小于0")
	}
	if rectwid < 0 {
		log.Fatal("宽度小于0")
	}
}


func main(){

	//思考为何Area和Diagonal首字母大写，自定义包中有解答
	mianji := rectangle.Area(rectlen,rectwid)
	duijiao := rectangle.Diagonal(rectlen,rectwid)

	fmt.Println(mianji,duijiao)
}
