package server

import (
	"go_SayHi/controllers/admin"
	"go_SayHi/controllers/api"
	"go_SayHi/pkg/config"
	"log/slog"
	"os"
	"strings"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/mlogclub/simple/web"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func NewServer() {
	conf := config.Instance

	app := iris.New()
	app.Logger().SetLevel("info")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   conf.AllowedOrigins,
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
	}))
	app.AllowMethods(iris.MethodOptions)

	app.OnAnyErrorCode(func(ctx iris.Context) {
		path := ctx.Path()
		var err error
		if strings.Contains(path, "/api/admin/") {
			err = ctx.JSON(web.JsonErrorCode(ctx.GetStatusCode(), "http error"))
		}
		if err != nil {
			slog.Error(err.Error(), slog.Any("err", err))
		}
	})

	app.Any("/", func(i iris.Context) {
		_ = i.JSON(map[string]interface{}{
			"engine": "go_SayHi",
		})
	})

	app.HandleDir("/admin", "./admin")

	// api
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/user").Handle(new(api.UserController))
	})

	// admin
	mvc.Configure(app.Party("/api/admin"), func(m *mvc.Application) {
		m.Party("/user").Handle(new(admin.UserController))
	})

	if err := app.Listen(":"+conf.Port,
		iris.WithConfiguration(iris.Configuration{
			DisableStartupLog:                 false,
			DisableInterruptHandler:           false,
			DisablePathCorrection:             false,
			EnablePathEscape:                  false,
			FireMethodNotAllowed:              false,
			DisableBodyConsumptionOnUnmarshal: false,
			DisableAutoFireStatusCode:         false,
			EnableOptimizations:               true,
			TimeFormat:                        "2006-01-01",
			Charset:                           "UTF-8",
		})); err != nil {
		slog.Error(err.Error(), slog.Any("err", err))
		os.Exit(-1)
	}
}
