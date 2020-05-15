package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Num   string `json:"num"`
	Pass  string `json:"pass"`
	Phone int    `json:"phone"`
}
type Jwt struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func Login(name string) (*User, error) {
	mods := &User{}
	err := DB.Get(mods, "select * from user where name=?", name)
	return mods, err
}
