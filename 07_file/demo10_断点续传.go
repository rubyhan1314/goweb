package main

import (
	"fmt"
	"os"
	"strconv"
	"io"
)

func main() {
	/*
	断点续传：
		文件传递：文件复制
			/Users/ruby/Documents/pro/aa.jpeg

		复制到
			/Users/ruby/Desktop/aa.jpeg

	思路：
		边复制，边记录复制的总量
	 */

	srcFile:="/Users/ruby/Documents/pro/aa.jpeg"
	destFile:="/Users/ruby/Desktop/aa.jpeg"
	tempFile:=destFile+"temp.txt"
	//fmt.Println(tempFile)
	file1,_:=os.Open(srcFile)
	file2,_:=os.OpenFile(destFile,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	file3,_:=os.OpenFile(tempFile,os.O_CREATE|os.O_RDWR,os.ModePerm)


	defer file1.Close()
	defer file2.Close()
	//1.读取临时文件中的数据，根据seek
	file3.Seek(0,0)
	bs:=make([]byte,100,100)
	n1,err:=file3.Read(bs)
	fmt.Println(n1)
	countStr:=string(bs[:n1])
	fmt.Println(countStr)
	//count,_:=strconv.Atoi(countStr)
	count,_:=strconv.ParseInt(countStr,10,64)
	fmt.Println(count)

	//2. 设置读，写的偏移量
	file1.Seek(count,0)
	file2.Seek(count,0)
	data:=make([]byte,1024,1024)
	n2:=-1// 读取的数据量
	n3:=-1//写出的数据量
	total :=int(count)//读取的总量

	for{
		//3.读取数据
		n2,err=file1.Read(data)
		if err ==io.EOF{
			fmt.Println("文件复制完毕。。")
			file3.Close()
			os.Remove(tempFile)
			break
		}
		//将数据写入到目标文件
		n3,_=file2.Write(data[:n2])
		total += n3
		//将复制总量，存储到临时文件中
		file3.Seek(0,0)
		file3.WriteString(strconv.Itoa(total))


		//假装断电
		//if total>8000{
		//	panic("假装断电了。。。，假装的。。。")
		//}
	}

}
