package router
 
import (
	"net/http"
	"io"
	"github.com/Luooojunnn/gono/utils"
	"encoding/json"
	"strconv"
	// "fmt"
	"net/url"
)

type routes struct {
	path string
}

func (this *routes) Then(cb func(ctx *utils.Ctx)) (*routes) {
	m := utils.MethodCbMap[this.path]
	if m == nil {
		// this's the first Thenfunc on this path
		// create new M instance
		cbs :=  []func(ctx *utils.Ctx) {
			cb,
		}
		m := &utils.M{
			Cbs: cbs,
		}
		utils.MethodCbMap[this.path] = m
	} else {
		m.Cbs = append(m.Cbs, cb)
	}
	return this
}


// HTTP Method: GET
func G(path string) (*routes) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			despatchMethod(path, w, r)
		} else {
			io.WriteString(w, "Error API Method " + path)
		}
	})
	return &routes{path: path} 
}

// HTTP Method: POST
func P(path string) (*routes) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			despatchMethod(path, w, r)
		} else {
			io.WriteString(w, "Error API Method " + path)
		}
	})
	return &routes{path: path} 
}

// despatch method
func despatchMethod(path string, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	query := r.Form
	ctx := &utils.Ctx{}
	ctx = mount(ctx, "Query" , query)
	cbs := utils.MethodCbMap[path].Cbs
	for i := 0; i < len(cbs); i++ {
		cbs[i](ctx)
	}
	switch v := ctx.Body.(type) {
	case string:
		io.WriteString(w, v)
	case int:
		io.WriteString(w, strconv.Itoa(v))
	case map[string]interface{}:
		w.Header().Set("Content-type", "application/json; charset=utf-8")
		jsn, _ := json.Marshal(v)
		io.WriteString(w, string(jsn))
	}
}

func mount(c *utils.Ctx, key string , value url.Values) (*utils.Ctx) {
	switch key {
	case "Query":
		c.Query = value
	}
	return c
}