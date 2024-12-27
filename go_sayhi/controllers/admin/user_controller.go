package admin

import (
	"fmt"
	"go_SayHi/models"
	"go_SayHi/services"
	"log/slog"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
)

type UserController struct {
	Ctx iris.Context
}

// func (c *UserController) GetSynccount() *web.JsonResult {
// 	go func() {
// 		services.UserService.SyncUserCount()
// 	}()
// 	return web.JsonSuccess()
// }

func (c *UserController) GetBy(id int64) *web.JsonResult {
	t := services.UserService.Get(id)
	fmt.Print(string(id))
	if t == nil {
		return web.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return web.JsonData(c.buildUserItem(t, true))
}

func (c *UserController) PostCreate() *web.JsonResult {
	username := c.Ctx.PostValue("username")
	nickname := c.Ctx.FormValue("nickname")
	password := c.Ctx.FormValue("password")
	rePassword := c.Ctx.FormValue("rePassword")
	email := c.Ctx.FormValue("email")

	slog.Info(username + nickname + password + rePassword + email)
	t, err := services.UserService.SignUp(username, email, nickname, password, rePassword)

	if err != nil {
		return web.JsonErrorMsg(err.Error())
	}

	return web.JsonData(c.buildUserItem(t, true))
}

func (c *UserController) buildUserItem(user *models.User, buildRoleIds bool) map[string]interface{} {
	b := web.NewRspBuilder(user).
		Put("roles", user.GetRoles()).
		Put("username", user.Username.String).
		Put("email", user.Email.String).
		Put("score", user.Score).
		Put("forbidden", user.IsForbidden())
	if buildRoleIds {
		// b.Put("roleIds", services.UserRoleService.GetUserRoleIds(user.Id))
	}
	return b.Build()
}
