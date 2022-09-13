package tokens

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//生成token与验证

type jwtClaims struct {
	jwt.StandardClaims     //jwt-go包预定义的一些字段
	Id                 int `json:"id"`
}

var (
	key        = "miniProject"
	ExpireTime = 604800 //token过期时间
)

//我自己往token里写进去的只有id
func GenerateToken(id int) string {
	claims := &jwtClaims{
		Id: id,
	}
	fmt.Println("1", id)
	//签发者和过期时间
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(*claims)
	if err != nil {
		log.Print("produceToken err:")
		fmt.Println(err)
		return ""
	}
	return singedToken
}

func genToken(claims jwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return singedToken, nil
}

//验证token
func VerifyToken(token string) (int, error) {
	TempToken, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return 0, errors.New("token解析失败")
	}
	claims, ok := TempToken.Claims.(*jwtClaims)
	if !ok {
		return 0, errors.New("发生错误")
	}
	if err := TempToken.Claims.Valid(); err != nil {
		return 0, errors.New("发生错误")
	}
	return claims.Id, nil
}
