package main

import (
	"time"

	"golang.cafe/pkg/option/server"
)

func main() {
	svr := server.New(
		server.WithHost(host:	"localhost"),
		server.WithMaxConn(maxConn: 100),
		server.WithPort(Port 8080),
		server.WithTimeout(Timeout: time.Minute),
	)
	svr.Start()
	time.Sleep(time.Second)
	svr.Stop()
}
