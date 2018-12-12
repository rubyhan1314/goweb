package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	/*
	读取文件：
		1.打开文件
		2.读取文件
			file.Read([]byte)-->n,err
			从文件中开始读取数据，存入到byte切片中，返回值n是本次实际读取的数据量
				如果读取到文件末尾，n为0,err为EOF：end of file
		3.关闭文件
	 */
	//step1：打开文件
	fileName := "/Users/ruby/Documents/pro/aa.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件有误：", err.Error())
		return
	}
	//step2：读/写
	bs := make([] byte, 4, 4)
	n := -1
	for{
		n,err=file.Read(bs)
		if n == 0 || err == io.EOF{
			fmt.Println("读取到文件末尾了，结束读取操作。。")
			break
		}
		fmt.Println(string(bs[:n]))
	}




	/*
	//第一次读取
	n, err := file.Read(bs)//从file对应的文件中读取最多len(bs)个数据，存入到bs数组中，n是实际读入的数量
	fmt.Println(err)//nil
	fmt.Println(n)//4
	fmt.Println(bs)//[97 98 99 100]
	fmt.Println(string(bs))//abcd
	//二次读取
	n ,err = file.Read(bs)

	fmt.Println(err)//nil
	fmt.Println(n)//4
	fmt.Println(bs)//[101 102 103 104]
	fmt.Println(string(bs))//efgh

	//三次读取
	n ,err = file.Read(bs)

	fmt.Println(err)//nil
	fmt.Println(n)//3
	fmt.Println(bs)//[105 106 10 104]
	fmt.Println(string(bs[:n]))//ij

	//第四次,读到末尾了
	n,err = file.Read(bs)
	fmt.Println(err)//EOF:end of file
	fmt.Println(n)//0
*/

	//step3：关闭文件
	file.Close()
}
/*
练习1：读取本地的文本文件
练习2：读取本地的图片文件，不用控制台打印，统计读取的数量即可。
练习3：将程序中的数据，保存到文件中。
 */
