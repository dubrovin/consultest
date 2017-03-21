package main

import (
	"fmt"
	"time"

	"bitbucket.infotech.team/gol/GoConsul"
	"github.com/dubrovin/consultest/server"
)

func main() {
	consul, err := GoConsul.NewConsulClient("10.0.1.201:8500")
	if err != nil {
		fmt.Println(err)
	}
	go server.Run(":8000")
	time.Sleep(time.Second)
	entry, _, err := consul.Service("app1", "default")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v %v", entry[0].Service.Address, entry[0].Service.Port)
	time.Sleep(time.Second * 60)

}
