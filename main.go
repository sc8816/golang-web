package main

import (
	"awesomeProject/control"
	"awesomeProject/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func mid(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		token := r.Form.Get("token")
		if token == "" {
			fmt.Println("无效token")
			return
		}
		d := &model.Jwt{}
		j, err := jwt.ParseWithClaims(token, d, func(token *jwt.Token) (i interface{}, e error) {
			return []byte("token"), nil
		})
		if err != nil {
			fmt.Println("token非法", err.Error())
		}
		//合法
		if j.Valid {
			next.ServeHTTP(w, r)
		}
	})
}
func main() {
	http.Handle("/list", mid(http.HandlerFunc(control.ClassList)))
	http.HandleFunc(`/login.html`, control.Login)
	http.HandleFunc(`/api/class/get`, control.ClassGet)
	http.HandleFunc(`/api/class/edit`, control.ClassEdit)
	http.HandleFunc(`/upload`, control.Upload)
	http.HandleFunc(`/api/class/delete`, control.ClassDelete)
	http.HandleFunc(`/api/class/list`, control.ClassList)
	http.HandleFunc(`/api/class/add`, control.ClassAdd)
	fmt.Println("is running")
	http.ListenAndServe(`:80`, nil)
}
