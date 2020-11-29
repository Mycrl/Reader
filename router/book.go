package router

import (
    "Reader/disk"
	"Reader/storage"
    "Reader/command"
    "strconv"
    "path"
	"github.com/kataras/iris/v12"
)

func Book(ctx iris.Context) {
	model := ctx.Values().Get("Model").(*storage.Model)
	name := ctx.Params().GetString("name")
	chapters, e := model.Query(name)
	AssertInternalServerError(e)
	ctx.JSON(chapters)
}

func Chapter(ctx iris.Context) {
    conf := ctx.Values().Get("Conf").(*command.Configure)
    name := ctx.Params().GetString("name")
    offset, o_e := strconv.Atoi(ctx.URLParam("offset"))
    AssertInternalServerError(o_e)
    size, s_e := strconv.Atoi(ctx.URLParam("size"))
    AssertInternalServerError(s_e)
    p := path.Join(conf.AppData, name)
    buf, e := disk.Reader(p, int64(offset), int64(size))
    AssertInternalServerError(e)
    ctx.Write(buf)
}