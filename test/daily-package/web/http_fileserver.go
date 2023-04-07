package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	FilePathName string
)

func init() {
	flag.StringVar(&FilePathName,"fpn","./","the file name")
}

func main() {

	flag.Parse()

	mux := http.NewServeMux()

	// 将地址前缀 /static 替换为 /public
	mux.Handle("/static/",http.StripPrefix("/static",http.FileServer(http.Dir("./public"))))
	//mux.Handle("/static/",http.FileServer(http.Dir("")))
	// 通过 http.ServeContent 直接输出文件内容
	mux.HandleFunc("/get",FileContentHandle)
	// 通过 命令行输入 参数 下载文件
	mux.Handle("/download/",http.StripPrefix("/download/",http.FileServer(http.Dir(FilePathName))))

	// mux.HandleFunc("/test/",testPath)

	server := http.Server{Addr: ":8080",Handler: mux}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// FileContentHandle 输出文件内容
func FileContentHandle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fileName := query.Get("show")
	if fileName == "" {
		w.WriteHeader(400)
		fmt.Fprintln(w,"filename is empty")
		return
	}
	GetFileContent(w,r,fileName,time.Time{})
}

// GetFileContent 读取文件内容函数
func GetFileContent(w http.ResponseWriter, r *http.Request,name string, time time.Time) {
	// 打开文件获取句柄
	f,err := os.Open(name)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w,"open file error, ", err)
		return
	}
	defer f.Close()

	fs,err := f.Stat()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w,"get file message failed, ",err)
		return
	}

	if fs.IsDir() {
		w.WriteHeader(500)
		fmt.Fprintln(w,"no such file ")
		return
	}

	http.ServeContent(w,r,name,fs.ModTime(),f)
}

func testPath (w http.ResponseWriter, r *http.Request) {
	// 通过 http.DetectContentType 获取返回的 Content-type 类型
	path := "/static/dir1/img.png"
	ext := filepath.Ext(path)
	fmt.Fprintln(w,ext)
	w.WriteHeader(http.StatusOK)
}
