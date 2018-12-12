package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	/*
	复制文件：
	 */
	//srcFile := "/Users/ruby/Documents/pro/aa.txt"
	//destFile := "/Users/ruby/Documents/aa.txt"

	srcFile :="/Users/ruby/Documents/pro/aa.jpeg"
	destFile:="image/bb.jpeg"
	total,err:=copyFile2(srcFile,destFile)
	fmt.Println(err)
	fmt.Println(total)

}
/*
该函数的功能：实现文件的拷贝，返回值是拷贝的总数量(字节),错误
 */
func copyFile1(srcFile,destFile string)(int,error){
	file1,err:=os.Open(srcFile)
	if err != nil{
		return 0,err
	}
	file2,err:=os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err !=nil{
		return 0,err
	}
	defer file1.Close()
	defer file2.Close()
	//拷贝数据
	bs := make([]byte,1024,1024)
	n :=-1//读取的数据量
	total := 0
	for {
		n,err = file1.Read(bs)
		if err == io.EOF || n == 0{
			fmt.Println("拷贝完毕。。")
			break
		}else if err !=nil{
			fmt.Println("报错了。。。")
			return total,err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total,nil

}

func copyFile2(srcFile, destFile string)(int64,error){
	file1,err:=os.Open(srcFile)
	if err != nil{
		return 0,err
	}
	file2,err:=os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err !=nil{
		return 0,err
	}
	defer file1.Close()
	defer file2.Close()

	return io.Copy(file2,file1)
}

