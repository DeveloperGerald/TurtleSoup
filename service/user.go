package service

import (
	"fmt"

	"github.com/DeveloperGerald/TurtleSoup/model"
	"github.com/DeveloperGerald/TurtleSoup/pkg/jwt"
	"github.com/DeveloperGerald/TurtleSoup/pkg/util"
	"github.com/DeveloperGerald/TurtleSoup/repository"
)

func RegisterUser(name, password string, email, phone *string) (*model.User, error) {
	// 加密密码
	hashedPassword, err := util.EncryptPassword(password)
	if err != nil {
		return nil, err
	}

	// 创建用户对象
	user := model.User{
		Name:         name,
		NickName:     name,
		PasswordHash: hashedPassword,
		Email:        email,
		Phone:        phone,
	}

	// 存储用户到数据库
	created, err := repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func LoginUser(name, password string) (string, error) {
	// 根据用户名查询用户
	user, err := repository.GetUserByName(name)
	if err != nil {
		return "", err
	}

	// 验证密码
	if !util.ComparePassword(user.PasswordHash, password) {
		return "", fmt.Errorf("密码错误")
	}

	// 密码正确，生成 JWT
	token, err := jwt.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
