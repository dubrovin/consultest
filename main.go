package main

import (
	"github.com/dubrovin/consultest/server"
	"bitbucket.infotech.team/gol/goconsul"
	"fmt"
	"time"
	"os"
)

func main() {
	consul, err := GoConsul.NewConsulClient(os.Getenv("CONSUL_ADDR"))
	if err != nil {
		fmt.Println(err)
	}
	go server.Run()
	time.Sleep(time.Second)
	_, _, err = consul.Service(os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_TAG"))

	if err != nil {
		fmt.Println(err)
	}

}
