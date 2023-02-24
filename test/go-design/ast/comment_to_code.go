package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"text/template"
)

const suffix = "_msg_gen.go"

const tpl = `

// code generate by github.com/ShenTong-2050/gen-const-msg DO NOT EDIT

// {{.pkg}} const code comment msg
package {{.pkg}}

// noErrorMsg if code is not found, GetMsg will return this
const noErrorMsg = "unknown error"

// messages get msg from const comment
var messages = map[int]string{
	{{range $key, $value := .comments}}
	{{$key}}: "{{$value}}",{{end}}
}

// GetMsg get error msg
func GetMsg(code int) string {
	var (
		msg string
		ok  bool
	)
	if msg, ok = messages[code]; !ok {
		msg = noErrorMsg
	}
	return msg
}
`

// gen 生成代码
func gen(comments map[string]string) ([]byte, error) {
	var buf = bytes.NewBufferString("")

	data := map[string]interface{}{
		"pkg":      os.Getenv("GOPACKAGE"),
		"comments": comments,
	}

	t, err := template.New("").Parse(tpl)
	if err != nil {
		return nil, fmt.Errorf("template init err %v",err)
	}

	err = t.Execute(buf, data)
	if err != nil {
		return nil, fmt.Errorf("template data err %v",err)
	}

	return format.Source(buf.Bytes())
}

func main() {

	file := os.Getenv("GOFILE")

	// 保存注释信息
	var comments = make(map[string]string)

	fSet := token.NewFileSet()

	f,err := parser.ParseFile(fSet,file,nil,parser.ParseComments)

	if err != nil {
		panic(err)
	}

	// 返回的是 key => Node 结构体、val => 多行注释结构体 的 map
	cmap := ast.NewCommentMap(fSet,f,f.Comments)

	for node := range cmap {
		// 仅支持 一条声明语句 一个 常量 情况
		if spec,ok := node.(*ast.ValueSpec); ok && len(spec.Names) == 1 {
			// 提取常量 注释
			ident := spec.Names[0]
			if ident.Obj.Kind == ast.Con {
				// 获取注释信息
				comments[ident.Name] = getComment(ident.Name,spec.Doc)
			}
		}
	}
}

// getComment 获取注释信息,来自 AST 标准库的 summary 方法
func getComment(name string,groups *ast.CommentGroup) string {

	var buf bytes.Buffer

	for _,comment := range groups.List {
		// 注释信息 会以 // 参数名，开始 我们实际使用时不需要，去掉
		text := strings.TrimSpace(strings.TrimPrefix(comment.Text,fmt.Sprintf("// %s",name)))
		buf.WriteString(text)
	}

	bytes := buf.Bytes()
	for i,b := range bytes {
		switch b {
		case '\t','\n','\r':
			bytes[i] = ' '
		}
	}

	return string(bytes)
}
