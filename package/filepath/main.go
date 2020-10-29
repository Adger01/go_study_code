package main
import (
	"fmt"
	"os"
	"path/filepath"
)

func example(file string){
	fmt.Println("file:",file)

	fmt.Printf("dir:%s \n",filepath.Dir(file))

	//检测地址是否是绝对地址，是绝对地址直接返回，不是绝对地址，会添加当前工作路径到参数path前，然后返回
	absStr,_ := filepath.Abs(file)
	fmt.Printf("abs:%s \n",absStr)

	//
	fmt.Println("base:",filepath.Base(file))

	//合并path
	fmt.Println(filepath.Join("/f","b","s"))

	//ext
	fmt.Println("filepath ext:",filepath.Ext(file))

	//isabs
	fmt.Println("IsAbs:",filepath.IsAbs(file))

	//os.PathSeparator == filepath.Separator
	fmt.Printf("PathSeparator:%c \n",os.PathSeparator)
	fmt.Printf("filepath.Separator:%c \n",filepath.Separator)

	// 清理路径中的多余字符，比如 /// 或 ../ 或 ./
	fmt.Println("filepath.clean:",filepath.Clean(file))

	//将相关平台的路径分隔符转为/
	fmt.Println("s:",filepath.ToSlash(file))

	//把路径分为dir和basefile
	d,f := filepath.Split(file)
	fmt.Println(d)
	fmt.Println(f)


	//
	matched,e := filepath.Match("*","/a/b")
	fmt.Println(matched)
	fmt.Println(e)
}

func main()  {
	file := "/Users/lijie//go/src/lijie.io/go_study_code/package/filepath/a.txt"
	//file := "./a.txt"
	fmt.Println(file)
	example(file)
}
