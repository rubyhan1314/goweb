package main

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	/*
	ioutil包：
		ReadFile()
		WriteFile()
		ReadDir()
		..
	 */

	//1.读取文件中的所有的数据
	//fileName1 := "/Users/ruby/Documents/pro/aa.txt"
	//data, err := ioutil.ReadFile(fileName1)
	//fmt.Println(err)
	//fmt.Println(string(data))

	//2.写出数据
	//fileName2:="/Users/ruby/Documents/pro/bbb.txt"
	//s1:="helloworld面朝大海春暖花开"
	//err:=ioutil.WriteFile(fileName2,[]byte(s1),0777)
	//fmt.Println(err)

	//3.
	s2:="qwertyuiopsdfghjklzxcvbnm"
	r1:=strings.NewReader(s2)
	data,_:=ioutil.ReadAll(r1)
	fmt.Println(data)

	//4.ReadDir(),读取一个目录下的子内容：子文件和子目录，但是仅有一层
	dirName:="/Users/ruby/Documents/pro"
	fileInfos,_:=ioutil.ReadDir(dirName)
	fmt.Println(len(fileInfos))
	for i:=0;i<len(fileInfos);i++{
		//fmt.Printf("%T\n",fileInfos[i])
		fmt.Println(i,fileInfos[i].Name(),fileInfos[i].IsDir())

	}

}
