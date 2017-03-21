package main

import (
	"github.com/dubrovin/consultest/server"
	"bitbucket.infotech.team/gol/goconsul"
	"fmt"
	"time"
)

func main() {
	consul, err := GoConsul.NewConsulClient("consul.service.consul:8500")
	if err != nil {
		fmt.Println(err)
	}
	go server.Run()
	time.Sleep(time.Second)
	_, _, err = consul.Service("Hello", "")

	if err != nil {
		fmt.Println(err)
	}

}
