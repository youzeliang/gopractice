package main

import (
	"crypto/md5"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
)

func main() {
	//m, _ := ParseFile("./errcode.go")
	//fmt.Printf("%#v", m)

	fmt.Println(GetSign("71118074", "180150386", "1746357"))
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func GetSign(stuId, stuCouId, planId string) (token string) {
	sign := "Xes-Report-Token-" + stuId + "-" + stuCouId + "-" + planId
	return Md5(sign)
}

func ParseFile(fileName string) (map[string]string, error) {
	fset := token.NewFileSet()
	// 解析文件，主要解析token
	f, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// 类型检查, 得到常量的值
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	for _, v := range f.Scope.Objects {
		if v.Kind == ast.Con {
			d := v.Decl.(*ast.ValueSpec)
			m[pkg.Scope().Lookup(v.Name).(*types.Const).Val().String()] = d.Comment.Text()
		}
	}
	return m, nil
}
