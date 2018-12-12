package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlystudents struct {
	XMLName xml.Name `xml:"students"`
	Version string `xml:"version,attr"`
	Students []student `xml:"student"`
	Description string `xml:",innerxml"`
}
type student struct {
	XMLName xml.Name `xml:"student"`
	StudentName string `xml:"studentName"`
	Age int `xml:"age"`
	Sex string `xml:"sex"`
	Books Recurlybookss `xml:"books"`
}


type Recurlybookss struct {
	XMLName xml.Name `xml:"books"`
	Version string `xml:"version,attr"`
	Books []book `xml:"book"`
	Description string `xml:",innerxml"`
}

type book struct {
	XMLName xml.Name `xml:"book"`
	BookName  string `xml:"bookName"`
	Price  string `xml:"price"`
}

func main() {
	file, err := os.Open("students.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlystudents{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}
