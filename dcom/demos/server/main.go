package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/gophertrain/material/dcom/demos/server/stringsvc"
)

func main() {
	upper := new(stringsvc.Upper)
	rpc.Register(upper)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)

}
