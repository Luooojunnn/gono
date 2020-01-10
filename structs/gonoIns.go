package structs

import (
	"fmt"
	"strconv"
	"net/http"
	"log"
)

type GonoIns struct {
	// Port int
}

func Init() (GonoIns)  {
	return GonoIns{}
}

func (this *GonoIns) Listen(port int)  {
	defer fmt.Printf("server is lisenning localhost:%d \n\n", port)
	err := http.ListenAndServe("localhost:" + strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	} 
}

// func (this *GonoIns) Use(arg interface{}) {
// 	fmt.Println(arg)
// }