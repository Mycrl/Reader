package router

import (
	"Reader/storage"
    "Reader/command"
    "io/ioutil"
	"github.com/kataras/iris/v12"
)

func AssertInternalServerError(err error) {
	if err != nil {
		panic(err)
	}
}

func Middleware(model *storage.Model, conf *command.Configure) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		ctx.Values().Set("Model", model)
        ctx.Values().Set("Conf", conf)
		ctx.Next()
	}
}

func InternalServerError(ctx iris.Context) {
	ctx.StatusCode(400)
	ctx.Writef("")
}

func Index(ctx iris.Context) {
    bytes, err := ioutil.ReadFile("./view/dist/index.html")
    AssertInternalServerError(err)
    ctx.Write(bytes)
}