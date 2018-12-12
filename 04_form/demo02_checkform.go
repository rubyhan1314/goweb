package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"regexp"
)


func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//1.验证必填字段
	//username := r.Form["username"][0]
	username := r.Form.Get("username")
	if len(username) == 0 {
		fmt.Println("用户名不能为空！")
		fmt.Fprintf(w, "用户名不能为空！") //这个写入到w的是输出到客户端的
	}
	//2.验证数字
	age, err := strconv.Atoi(r.Form.Get("age"))
	if err != nil {
		//数字转化出错了，那么可能就是不是数字
		fmt.Println("您输入的不是数字！")
		fmt.Fprintf(w, "您输入的不是数字！") //这个写入到w的是输出到客户端的
	}
	//接下来就可以判断这个数字的大小范围了
	if age > 100 || age < 0 {
		//太大了或太小了
		fmt.Println("您输入的年龄太大了或太小了，请输入0-100之间的整数！")
		fmt.Fprintf(w, "您输入的年龄太大了或太小了，请输入0-100之间的整数！") //这个写入到w的是输出到客户端的
	}

	//或者正则表达式
	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		fmt.Println("验证有误，您输入的年龄太大了或太小了！")
		fmt.Fprintf(w, "验证有误，您输入的年龄太大了或太小了！")
	}

	//3.验证中文
	if m, _ := regexp.MatchString(`^[\x{4e00}-\x{9fa5}]+$`, r.Form.Get("zhname")); !m {
		fmt.Println("验证有误，请输入中文！")
		fmt.Fprintf(w, "验证有误，请输入中文！")
	}

	//4. 验证英文
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("enname")); !m {
		fmt.Println("验证有误，请输入英文！")
		fmt.Fprintf(w, "验证有误，请输入英文！")
	}

	//5. 邮箱
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
		fmt.Println("请输入正确邮箱地址")
		fmt.Fprintf(w, "验证有误，请输入正确邮箱地址！")
	}

	//6. 验证手机号
	if m, _ := regexp.MatchString(`^(1[3|5|6|7|8][0-9]\d{8})$`, r.Form.Get("mobile")); !m {
		fmt.Println("请输入正确手机号码")
		fmt.Fprintf(w, "验证有误，请输入正确手机号码！")
	}

	//7. 下拉菜单
	xueli := r.Form.Get("xueli")
	res1 := checkSelect(xueli)
	if !res1 {
		fmt.Println("请选择正确的下拉列表！")
		fmt.Fprintf(w, "请选择正确的下拉列表！")
	}

	// 8. 单选按钮
	sex := r.Form.Get("sex")
	res2 := checkSex(sex)
	if !res2 {
		fmt.Println("请选择正确的性别！")
		fmt.Fprintf(w, "请选择正确的性别！")
	}

	// 9. 复选框
	hobby := r.Form["hobby"]
	res3 := checkHobby(hobby)
	if !res3 {
		fmt.Println("请选择正确的爱好！")
		fmt.Fprintf(w, "请选择正确的爱好！")
	}


	// 10 身份证号
	//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
		fmt.Println("请选择正确的身份证号！")
		fmt.Fprintf(w, "请选择正确的身份证号！")
	}


	//fmt.Println("验证成功！")
	//fmt.Fprintf(w, "验证成功！")

}

/**
验证下拉列表
 */
func checkSelect(xueli string) bool {
	slice := []string{"xiaoxue", "chuzhong", "gaozhong", "dazhuan", "benke", "shuoshi", "boshi", "lieshi"}
	for _, v := range slice {

		if v == xueli {
			return true
		}
	}
	return false
}

/**
验证单选按钮
 */
func checkSex(sex string) bool {
	slice := []string{"male", "female", "other"}
	for _, v := range slice {
		if v == sex {
			return true
		}
	}
	return false
}

/**
验证复选框
 */
func checkHobby(hobby []string) bool {
	slice := []string{"game", "girl", "money", "power"}

	hobby2 := Slice_diff(hobby, slice)

	if hobby2 == nil {
		return true
	}
	return false
}


func Slice_diff(slice1, slice2 []string) (diffslice []string) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

/**
判断是一个切片中是否包含指定的数值
 */
func InSlice(val string, slice []string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	http.HandleFunc("/register", register)   //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
