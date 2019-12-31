package structs

import (
	"fmt"
	"strconv"
	"net/http"
	_ "gono/router"
)

type GonoIns struct {
	Port int
}


// func (this *GonoIns) Use()  {
// 	fmt.Print(this)
// }

func Init() (GonoIns)  {
	return GonoIns{}
}

func (this *GonoIns) Listen(port int)  {
	fmt.Printf("server is lisenning localhost:%d \n\n", port)
	http.ListenAndServe("localhost:" + strconv.Itoa(port), nil)
}