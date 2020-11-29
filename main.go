package main

import (
	"Reader/disk"
	"Reader/router"
	"Reader/storage"
    "Reader/command"
	"github.com/kataras/iris/v12"
)

func Cors(ctx iris.Context) {
    ctx.Header("Access-Control-Allow-Origin", "*")
    ctx.Header("Access-Control-Allow-Methods", "*")
    ctx.Header("Access-Control-Allow-Headers", "*")
    ctx.Next()
}

func main() {
    conf := command.Configure{}
    (&conf).Parse()
	app := iris.Default()
    app.Use(Cors)
	model := storage.NewModel(conf.Db)
	go disk.NewWatch(conf.AppData, model).Run()
	app.UseGlobal(router.Middleware(model, &conf))
    app.HandleDir("/static", "./view/dist/static")
	app.OnErrorCode(iris.StatusInternalServerError, router.InternalServerError)
    app.Get("/chapter/{name:string}", router.Chapter)
	app.Get("/book/{name:string}", router.Book)
	app.Get("/list", router.List)
    app.Get("/", router.Index)
	app.Listen(":" + conf.Port)
}
