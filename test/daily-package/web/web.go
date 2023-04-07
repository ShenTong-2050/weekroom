package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

var ExtToContentType map[string]string = map[string]string{
	".html":"text/html;charset=utf-8",
	".js":"application/javascript",
	".css":"text/css;charset=utf-8",
	".xml":"text/xml;charset=utf-8",
	".jpg":"image/jpeg",
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/static/",FileHandle)
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,"hello world!")
	})
	server := &http.Server{Addr: ":8080",Handler: mux}
	if err := server.ListenAndServe(); err != nil {
		toHttpError(err)
	}
}

// FileHandle 文件处理程序
func FileHandle(w http.ResponseWriter,r *http.Request) {
	// 通过 url 获取请求地址 . 的作用是指相对路径
	path := "."+r.URL.Path

	// 打开文件
	f,err := os.Open(path)
	if err != nil {
		Error(w,toHttpError(err))
		return
	}
	defer f.Close()

	// 获取文件信息
	stat,err := f.Stat()
	if err != nil {
		Error(w,toHttpError(err))
		return
	}

	// 判断 是否为目录 是则遍历
	if stat.IsDir() {
		DirList(w,r,f)
	}

	// 读取所有文件内容
	data,err := ioutil.ReadAll(f)
	if err != nil {
		Error(w,toHttpError(err))
		return
	}

	// 根据文件后缀设置对应的文件响应类型 与 文件长度
	ext := filepath.Ext(path)
	// if contentType := ExtToContentType[ext]; contentType != "" {
	if contentType := http.DetectContentType([]byte(ext)); contentType != "" {
		w.Header().Set("Content-type",contentType)
	}
	w.Header().Set("Content-length",strconv.FormatInt(stat.Size(),10))
	w.Write(data)
}

// DirList 遍历目录
func DirList(w http.ResponseWriter,r *http.Request,f http.File) {
	dirs,err := f.Readdir(-1)
	if err != nil {
		Error(w,toHttpError(err))
		return
	}

	// 按照文件名称排序
	sort.Slice(dirs,func(i,j int) bool { return dirs[i].Name() < dirs[j].Name() })

	w.Header().Set("Content-type","charset=utf-8;charset=utf-8")

	fmt.Fprintf(w,"<pre>\n")
	for _,item := range dirs {
		// 获取 文件/目录 名称
		name := item.Name()
		if item.IsDir() {
			name+="/"
		}
		url := netUrl.URL{Path: name}
		fmt.Fprintf(w,"<a href=\"%s\">%s</a>\n",url.String(),name)
	}
	fmt.Fprintf(w,"</pre>\n")
}

func toHttpError(err error) int {
	if os.IsNotExist(err) {
		return http.StatusNotFound
	} else if os.IsPermission(err) {
		return http.StatusForbidden
	}
	return http.StatusInternalServerError
}

func Error(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}