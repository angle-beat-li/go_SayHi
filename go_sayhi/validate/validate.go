package validate

import (
	"errors"
	"regexp"
	"strings"

	"github.com/mlogclub/simple/common/strs"
)

// IsUsername 验证用户名合法性，用户名必须由5-12位（数字、字母、_、-）组成，且必须以字母开头
func IsUsername(userName string) error {
	if strs.IsBlank(userName) {
		return errors.New("用户名不能为空")
	}
	matched, err := regexp.MatchString("^[0-9a-zA-Z_-]{5,12}$", userName)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	matched, err = regexp.MatchString("^[a-zA-Z]", userName)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	return nil
}

// IsEmail 验证邮箱是否是合法邮箱
func IsEmail(email string) error {
	if strs.IsBlank(email) {
		return errors.New("邮箱不能为空")
	}
	matched, _ := regexp.MatchString(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`, email)
	if !matched {
		return errors.New("邮箱格式不符合规范")
	}
	return nil
}

// IsValidPassword 是否是合法的密码
func IsValidPassword(password, rePassword string) error {
	if err := IsPassword(password); err != nil {
		return err
	}
	if password != rePassword {
		return errors.New("两次输入的密码不一致")
	}
	return nil
}

func IsPassword(password string) error {
	if strs.IsBlank(password) {
		return errors.New("请输入密码")
	}
	if strs.RuneLen(password) < 6 {
		return errors.New("密码过于简单")
	}
	if strs.RuneLen(password) > 1024 {
		return errors.New("密码长度不能超过128")
	}
	return nil
}

// IsUrl 是否是合法的URL
func IsUrl(url string) error {
	if strs.IsBlank(url) {
		return errors.New("URL格式错误")
	}
	indexofHttp := strings.Index(url, "http://")
	indexofHttps := strings.Index(url, "https://")
	if indexofHttp == 0 || indexofHttps == 0 {
		return nil
	}
	return errors.New("url 格式错误")
}
