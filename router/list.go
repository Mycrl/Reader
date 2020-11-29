package router

import (
	"Reader/storage"
	"github.com/kataras/iris/v12"
)

func List(ctx iris.Context) {
	model := ctx.Values().Get("Model").(*storage.Model)
	list := model.List()
	ctx.JSON(list)
}
