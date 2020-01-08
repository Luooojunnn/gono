# gono

使用更 “前端” 一点的书写方式来写 go web 服务吧！


**1. 启动服务就是那么简单！**

❌
```golang
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

✅
```golang
package main

import (
	"gono"
	_ "./routes"
)

// var conf = structs.InitParams{
// 	Proxy: false,
// }

func main()  {
	gn := gono.Init()
	gn.Listen(9000)
}
```

**2. 请求拦截就是那么简单！**

✅
```golang
package main

import (
	
)
```
