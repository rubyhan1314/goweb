package main
import (
	"fmt"
	"os"
)
func main() {

	/*
	写出数据到文件：

	 */
	//file,err:=os.Open("/Users/ruby/Documents/pro/ab.txt")
	file,err:=os.OpenFile("/Users/ruby/Documents/pro/ab.txt",os.O_CREATE|os.O_WRONLY,os.ModePerm)
	if err != nil {
		fmt.Println("打开文件有误：", err.Error())
	}
	//defer file.Close()
	fmt.Println(file)
	bs := []byte{65, 66, 67, 68, 69, 70} //ABCDEF
	n, err := file.Write(bs)
	fmt.Println(err)
	fmt.Println(n)
	n, err = file.WriteString("hello面朝大海")
	fmt.Println(err)
	fmt.Println(n)

	file.Close()

	//file, _ = os.OpenFile("/Users/ruby/Documents/pro/ab.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)

	//file.WriteString("1234567890")

	//file.Close()
}
