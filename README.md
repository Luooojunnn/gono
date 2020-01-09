# gono

使用更 “前端” 一点的书写方式来写 go web 服务吧！


### 用法

**go version 1.11+**

> import "github.com/Luooojunnn/gono"

go run 你的项目的时候，自动在 go.sum 中下载保存 gono 信息啦

---

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

**GET**

```golang
package routes

import (
	"gono/router"
	"gono/utils"
	"fmt"
) 

// 模拟路由拆分

// get 方法
func init(){
	router.G("/").Then(func(ctx *utils.Ctx) {
		ctx.Body = 112233
		query := ctx.Query
		for key, val := range query {
			fmt.Println("key: " + key + ", value: ", val)
		}
	}).Then(func(ctx *utils.Ctx)  {
		fmt.Println("你好，世界")
	}).Then(func(ctx *utils.Ctx)  {
		fmt.Println("hello, world")
	})
}
```
**POST**

```golang
package routes

import (
	"gono/utils"
	"gono/router"
	"fmt"
) 

// 模拟路由拆分

func init()  {
	PostUser()
}
// post 方法
func PostUser() {
	router.P("/user").Then(func(ctx *utils.Ctx) {
		ctx.Body = map[string]interface{} {
			"name": "emine",
			"age": 20,
		}

		query := ctx.Query
		for key, val := range query {
			fmt.Println("key: " + key + ", value: ", val)
		}

		ctx.Extra = "我是要透传给下一个 then 的信息"
	}).Then(func(ctx *utils.Ctx) {
		fmt.Println(ctx.Extra)
	})
}

```

**3. 就是那么简单！**

目前来说，Then 里边的回调的入参 ctx 上边仅有：
 
 - Body - 类似 koa ，返回体的内容
 - Query - 请求携带的参数，在此处获取，是一个 map 结构
 - Extra - 额外信息，可以通过它将信息传递给下一个 Then

*请注意：使用链式的 Then ，入参函数会顺序执行，后边的 Then 会覆盖前边的赋值*
