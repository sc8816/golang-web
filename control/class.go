package control

import (
	"awesomeProject/model"
	"awesomeProject/utils"
	"fmt"
	"net/http"
	"strconv"
)

//列表 ClassLis
func ClassList(w http.ResponseWriter, r *http.Request) {
	mods, err := model.ListClass()
	if err != nil {
		w.Write(utils.FormatterResult(utils.Fail, "未查询到信息", err.Error()))
		return
	}
	w.Write(utils.FormatterResult(utils.Success, "请求成功", mods))
}

//删除ClassDelete
func ClassDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		fmt.Println("数据有误")
		return
	}
	if model.DelClass(id) {
		w.Write(utils.FormatterResult(utils.Success, "删除成功"))
		return
	}
	w.Write(utils.FormatterResult(utils.Fail, "删除失败"))
	return
}

//id查询
func ClassGet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		fmt.Println("数据有误")
		return
	}
	mod, err := model.GetClassById(id)
	if err == nil {
		w.Write(utils.FormatterResult(utils.Success, "查询成功", mod))
		return
	}
	w.Write(utils.FormatterResult(utils.Fail, "查询失败", err.Error()))
	return
}

//add
func ClassAdd(w http.ResponseWriter, r *http.Request) {
	mod := model.Class{
		Name: r.FormValue("name"),
		Desc: r.FormValue("desc"),
	}
	if mod.Name == "" {
		w.Write(utils.FormatterResult(utils.Fail, "名称不能为空"))
		return
	}
	id := model.AddClass(mod)
	if id != 0 {
		w.Write(utils.FormatterResult(utils.Success, "添加成功", id))
		return
	}
	w.Write(utils.FormatterResult(utils.Fail, "添加失败"))
}

//edit
func ClassEdit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.Write(utils.FormatterResult(utils.Fail, "id有误"))
		return
	}
	mod := model.Class{
		Id:   id,
		Name: r.FormValue("name"),
		Desc: r.FormValue("desc"),
	}
	if mod.Name == "" {
		w.Write(utils.FormatterResult(utils.Fail, "名称不能为空"))
		return
	}
	if model.EditClass(mod) {
		w.Write(utils.FormatterResult(utils.Success, "修改成功"))
		return
	}
	w.Write(utils.FormatterResult(utils.Fail, "修改失败"))
}
