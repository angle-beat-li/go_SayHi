package services

import (
	"errors"
	"go_SayHi/models"
	"go_SayHi/models/constants"
	"go_SayHi/repositories"
	"go_SayHi/validate"
	"strings"

	"github.com/mlogclub/simple/common/dates"
	"github.com/mlogclub/simple/common/passwd"
	"github.com/mlogclub/simple/common/strs"
	"github.com/mlogclub/simple/sqls"
)

// 邮箱验证邮件有效期
const emailVerifyExpireHour = 24

type userService struct {
}

func newUserService() *userService {
	return &userService{}
}

var UserService = newUserService()

func (s *userService) Get(id int64) *models.User {
	return repositories.UserRepository.Get(sqls.DB(), id)
}

func (s *userService) Take(where ...interface{}) *models.User {
	return repositories.UserRepository.Take(sqls.DB(), where...)
}

func (s *userService) Create(t *models.User) (err error) {
	err = repositories.UserRepository.Create(sqls.DB(), t)
	if err == nil {

	}
	return nil
}

// 用户注册
func (s *userService) SignUp(username, email, nickname, password, rePassword string) (*models.User, error) {
	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)
	nickname = strings.TrimSpace(nickname)
	password = strings.TrimSpace(password)
	rePassword = strings.TrimSpace(rePassword)

	// 验证昵称
	if len(nickname) == 0 {
		return nil, errors.New("昵称不能为空")
	}

	// 验证密码
	err := validate.IsValidPassword(password, rePassword)

	if err != nil {
		return nil, err
	}

	// 验证邮箱
	if err := validate.IsEmail(email); err != nil {
		return nil, err
	} else if s.isEmailExists(email) {
		return nil, errors.New("邮箱已经被使用")
	}

	// 验证用户名
	if err := validate.IsUsername(username); err != nil {
		return nil, err
	} else if s.isUsernameExists(username) {
		return nil, errors.New("该用户名已经被使用")
	}
	user := &models.User{
		Username:   sqls.SqlNullString(username),
		Email:      sqls.SqlNullString(email),
		Nickname:   nickname,
		Password:   passwd.EncodePassword(password),
		Status:     constants.StatusOk,
		CreateTime: dates.NowTimestamp(),
		UpdateTime: dates.NowTimestamp(),
	}
	return user, repositories.UserRepository.Create(sqls.DB(), user)
}

func (s *userService) GetByEmail(email string) *models.User {
	return repositories.UserRepository.GetByEmail(sqls.DB(), email)
}

// 判断邮箱是否存在
func (s *userService) isEmailExists(email string) bool {
	return repositories.UserRepository.GetByEmail(sqls.DB(), email) != nil
}

// 判断用户名是否存在
func (s *userService) isUsernameExists(username string) bool {
	return repositories.UserRepository.GetByUsername(sqls.DB(), username) != nil
}
func (s *userService) GetByUsername(username string) *models.User {
	return repositories.UserRepository.Take(sqls.DB(), "username = ?", username)
}

// 用户登录
func (s *userService) SingOn(username, password string) (*models.User, error) {
	if strs.IsBlank(username) {
		return nil, errors.New("请输入账号")
	}
	if strs.IsBlank(password) {
		return nil, errors.New("请输入密码")
	}
	if err := validate.IsPassword(password); err != nil {
		return nil, err
	}
	var user *models.User = nil
	if err := validate.IsEmail(username); err == nil {
		user = repositories.UserRepository.GetByEmail(sqls.DB(), username)
	} else {
		user = repositories.UserRepository.GetByUsername(sqls.DB(), username)
	}

	if user == nil || user.Status != constants.StatusOk {
		return nil, errors.New("用户名或者密码错误")
	}
	if !passwd.ValidatePassword(user.Password, password) {
		return nil, errors.New("用户名或者密码错误")
	}
	return user, nil
}

func (s *userService) UpdateColumn(userId int64, name string, value interface{}) error {
	return repositories.UserRepository.UpdateColumns(sqls.DB(), userId, name, value)
}

// 修改密码
func (s *userService) UpdatePassword(userId int64, oldPassword, newPassword, newRePassword string) error {
	if err := validate.IsValidPassword(newPassword, newRePassword); err != nil {
		return err
	}
	user := s.Get(userId)
	if len(user.Password) == 0 {
		errors.New("请先设置密码")
	}
	if !passwd.ValidatePassword(user.Password, oldPassword) {
		return errors.New("密码错误，请重试")
	}
	return s.UpdateColumn(userId, "password", passwd.EncodePassword(newPassword))
}
