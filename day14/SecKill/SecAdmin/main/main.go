package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "oldboy/day14/SecKill/SecAdmin/router"
)

func main() {
	err := initAll()
	if err != nil {
		panic(fmt.Sprintf("init database failed, err:%v", err))
		return
	}
	beego.Run()
}
