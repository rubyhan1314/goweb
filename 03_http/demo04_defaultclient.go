package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {
	request, err:=http.NewRequest(http.MethodGet,"http://www.chaindesk.cn/",nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	resp,err:=http.DefaultClient.Do(request)
	//resp,err := http.Get("http://www.chaindesk.cn/")
	//resp,err := http.Get("http://www.baidu.com/")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	s,err:=httputil.DumpResponse(resp,true)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",s)
}
