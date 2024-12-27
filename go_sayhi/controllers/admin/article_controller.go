package admin

import (
	"go_SayHi/pkg/sitemap"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
)

type ArticleController struct {
	Ctx iris.Context
}

func (c *ArticleController) GetSitemap() *web.JsonResult {
	go func() {
		sitemap.Generate()
	}()
	return web.JsonSuccess()
}
