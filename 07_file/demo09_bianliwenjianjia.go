package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	/*
	遍历文件夹
	 */
	dirName := "/Users/ruby/Documents/pro"
	listDir(dirName)
}

func listDir(dirName string){
	fileInfos,_:=ioutil.ReadDir(dirName)
	for _,fi:=range fileInfos{
		fileName := dirName+"/"+fi.Name()
		fmt.Println(fileName)
		if fi.IsDir(){
			listDir(fileName)
		}
	}
}
