package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {

	src := `package main
	func main() {
		println("hello world")
	}
	`

	// 新建一个 抽象语法树【ast】 文件集合
	fSet := token.NewFileSet()

	// parser.ParseFile() 会 解析 单个 Go 源文件的源代码 并 返回相应的 ast.File【抽象语法树的单个节点】 节点
	// 源文件 可以通过 传入 源文件 的 文件名/src 参数提供
	// 如果 src != nil 则 ParserFile 将从 src 中解析源代码，文件名 为 仅在记录位置信息时使用
	f,err := parser.ParseFile(fSet,"",src,0)

	if err != nil {
		panic(err)
	}

	ast.Print(fSet,f)
}
