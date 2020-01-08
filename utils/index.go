package utils

type Ctx struct {
	Body interface{}
	Query map[string][]string
	Extra interface{}
}

type M struct {
	Cbs []func(ctx *Ctx)
}

var MethodCbMap map[string]*M

func init(){
	// init methodCbMap
	MethodCbMap = make(map[string]*M)
}