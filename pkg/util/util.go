package util

import (
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword 用于加密用户密码
func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword 用于验证密码是否与存储的密码匹配
func ComparePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func StructToMap(v interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// 反射获取结构体的值和类型
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return result
	}

	// 遍历结构体字段
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		// 获取结构体标签
		tag := fieldType.Tag.Get("json")
		if tag == "" {
			continue
		}

		// 检查字段是否为nil或零值
		if field.Kind() == reflect.Ptr && field.IsNil() {
			// 如果为nil，跳过该字段
			continue
		}

		// 如果字段有效，将字段添加到map中
		result[tag] = field.Interface()
	}

	return result
}
