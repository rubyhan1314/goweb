package main

import (
	"regexp"
	"fmt"
)

func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func IsNum(num string)(b bool){
	if m, _ := regexp.MatchString("^[0-9]+$", num); m {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println(IsIP("192.168.1.100"))
	fmt.Println(IsNum("9527"))
}