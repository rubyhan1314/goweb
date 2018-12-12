package main

import (
	"fmt"
	"path/filepath"
	"path"
)

func main() {
	/*
	文件操作：
		1.路径：
			相对路径：relative
				aa.txt
				都是相当于当前的工程
			绝对路径：absolute
				/home/ruby/文档/pro/aa.txt

			.当前目录
			..上一层

		2.创建文件夹,如果文件夹存在，创建失败
			os.MKdir()
			os.MKdirAll()

		3.创建文件：如果文件存在，会覆盖
			os.Create()-->*file

		4.打开文件：
			os.Open(filename)
			os.OpenFile(filename,mode,perm)
		5.关闭文件：
			file.Close()
		6.删除：
			os.Remove()
			os.RemoveAll()
	 */

	//1.路径
	fileName1:="/Users/ruby/Documents/pro/aa.txt"
	fileName2:="aa.txt"
	fmt.Println(filepath.IsAbs(fileName1))//true
	fmt.Println(filepath.IsAbs(fileName2))//false
	fmt.Println(filepath.Abs(fileName1)) ///Users/ruby/Documents/pro/aa.txt
	fmt.Println(filepath.Abs(fileName2))///Users/ruby/go/src/08_file/aa.txt

	fmt.Println("获取父目录：",path.Join(fileName1,"..")) ///Users/ruby/Documents/pro

	//2.创建目录
	//os.MKdir()，仅创建一层
	//os.MKdirAll(),创建需要的多层
	//err:=os.Mkdir("/Users/ruby/Documents/pro/bb",os.ModePerm)
	//if err !=nil{
	//	fmt.Println("err-->",err.Error())
	//}else{
	//	fmt.Println("文件夹创建成功。。")
	//}

	//err:= os.MkdirAll("/Users/ruby/Documents/pro/cc/dd/ee",os.ModePerm)
	//if err !=nil{
	//	fmt.Println("err-->",err.Error())
	//}else{
	//	fmt.Println("文件夹创建成功。。")
	//}

	//3.创建文件,如果文件已经存在，会被覆盖。
	//file1,err:=os.Create("/Users/ruby/Documents/pro/bb/abc.txt")
	//if err!=nil{
	//	fmt.Println("err:",err.Error())
	//}
	//fmt.Println(file1)
	//
	//file2,_:=os.Create(fileName2)//创建一个相对路径的文件
	//fmt.Println(file2)

	//4.打开文件,让当前的程序和指定的文件建立了一个链接
	//file3,err:=os.Open(fileName1)
	//if err !=nil{
	//	fmt.Println("\t",err)
	//}
	//fmt.Println(file3)

	/*
	第一个参数：文件名称
	第二个参数：文件的打开方式
		O_RDONLY：只读模式(read-only)
		O_WRONLY：只写模式(write-only)
		O_RDWR：读写模式(read-write)
		O_APPEND：追加模式(append)
		O_CREATE：文件不存在就创建(create a new file if none exists.)
	第三个参数：文件的权限：文件不存在创建文件，需要指定权限
	 */
	//file4,err:=os.OpenFile(fileName1,os.O_WRONLY|os.O_CREATE,os.ModePerm)

	//5.关闭文件,程序和文件之间的链接断开。
	//file3.Close()

	//6.删除文件和文件夹：慎用，慎用，再慎用
	/*
	remove(),删除文件和空文件夹
	removeAll(),删除所有
	 */
	//err:=os.Remove("/Users/ruby/Documents/pro//bb")
	//fmt.Println(err)
	//err:=os.RemoveAll("/Users/ruby/Documents/pro/cc")
	//fmt.Println(err)
}
