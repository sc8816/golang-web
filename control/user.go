package control

import (
	"awesomeProject/model"
	"awesomeProject/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	user, err := model.Login(name)
	if err != nil {
		w.Write(utils.FormatterResult(utils.Fail, "用户不存在", err.Error()))
		return
	}
	pass := r.FormValue("pass")
	if user.Pass != pass {
		w.Write(utils.FormatterResult(utils.Fail, "密码错误"))
		return
	}
	fmt.Println(user)
	data := model.Jwt{
		Id:   user.Id,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
		},
	}
	ss := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	token, err := ss.SignedString([]byte(`token`))
	fmt.Println(token, err)
	if err != nil {
		w.Write(utils.FormatterResult(utils.Fail, "生成token失败", err.Error()))
		return
	}
	w.Write(utils.FormatterResult(utils.Success, "登陆成功", token))
}
