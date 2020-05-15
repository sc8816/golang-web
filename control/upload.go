package control

import (
	"awesomeProject/utils"
	"io"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1 << 20)
	f, h, err := r.FormFile("data")
	if err != nil {
		w.Write(utils.FormatterResult(utils.Fail, "上传失败", err.Error()))
		return
	}
	dst, _ := os.Create(h.Filename)
	io.Copy(dst, f)
	f.Close()
	dst.Close()
	w.Write(utils.FormatterResult(utils.Success, "上传成功", "***"))
}
