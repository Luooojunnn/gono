package gono

import (
	"fmt"
	"gono/structs"
)


func Init(arg ...structs.InitParams) (structs.GonoIns)  {
	initParamsFc(arg)
	return structs.Init()
}

func initParamsFc(arg []structs.InitParams){
	if len(arg) == 0 {
		fmt.Println("no config")
	} else if len(arg) > 1 {
		panic("error config")
	} else {
		if arg[0].Proxy {
			fmt.Println("open gono proxy")
		}
	}
}