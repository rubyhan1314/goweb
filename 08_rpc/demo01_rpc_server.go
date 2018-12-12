package main

import (
	"net/rpc"
	"net/http"
	"log"
)

type Params struct {
	Width, Height int
}
type Rect struct{}


func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func main() {
	rect := new(Rect)
	//注册一个rect服务
	rpc.Register(rect)
	//把服务处理绑定到http协议上
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
