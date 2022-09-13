package model

import (
	"github.com/ShiinaOrez/GoSecurity/security"
)


//不同时刻生成的哈希不一样,所以得用下面的函数比较
func GeneratePasswordHash(password string) string {
	return security.GeneratePasswordHash(password)
}

func CheckPassword(password, hashPwd string) bool {
	return security.CheckPasswordHash(password, hashPwd)
}
