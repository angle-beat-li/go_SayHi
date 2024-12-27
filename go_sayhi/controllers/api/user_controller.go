package api

import (
	"go_SayHi/models"
	"go_SayHi/services"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
)

type UserController struct {
	Ctx iris.Context
}

func (c *UserController) GetBy(id int64) *web.JsonResult {
	t := services.UserService.Get(id)
	if t == nil {
		return web.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return web.JsonData(c.buildUserItem(t, true))
}

func (c *UserController) PostUpdatePassword() *web.JsonResult {
	// TODO: 后续添加登陆判断
	var (
		userId, _     = c.Ctx.PostValueInt64("userId")
		oldPassword   = c.Ctx.PostValue("oldPassword")
		newPassword   = c.Ctx.PostValue("newPassword")
		newRePassword = c.Ctx.PostValue("newRePassword")
	)
	if err := services.UserService.UpdatePassword(userId, oldPassword, newPassword, newRePassword); err != nil {
		return web.JsonErrorMsg(err.Error())
	}
	return web.JsonSuccess()
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
