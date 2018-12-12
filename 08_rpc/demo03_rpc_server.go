package main

import (
	"net/rpc"
	"log"
	"net"
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


	tcpAddr, err :=net.ResolveTCPAddr("tcp",":2001")
	if err != nil {
		log.Fatal(err)
	}

	listener,err := net.ListenTCP("tcp",tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for{
		conn,err :=listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}


}
