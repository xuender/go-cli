package utils

import (
	"go/ast"
	"unicode"
)

type Source struct {
	Package string
	Funcs   []string
	Structs []*Struct
}

type Struct struct {
	Name       string
	FieldNames []string
	Fields     []*ast.Field
}

func (p *Struct) String() string {
	return p.Name
}

func NewStruct(name string, stu *ast.StructType) *Struct {
	ret := &Struct{Name: name, Fields: stu.Fields.List}
	ret.FieldNames = make([]string, len(ret.Fields))

	for i, field := range stu.Fields.List {
		if len(field.Names) == 1 {
			ret.FieldNames[i] = field.Names[0].Name
		} else if len(field.Names) > 1 {
			for _, name := range field.Names {
				ret.FieldNames[i] += name.String() + ","
			}
		}
	}

	return ret
}

func NewSource(files ...string) (*Source, error) {
	file, err := Parse(files...)
	if err != nil {
		return nil, err
	}

	source := &Source{Funcs: []string{}, Structs: []*Struct{}}

	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.TypeSpec:
			if stu, isStruct := node.Type.(*ast.StructType); isStruct {
				if unicode.IsUpper([]rune(node.Name.Name)[0]) {
					source.Structs = append(source.Structs, NewStruct(node.Name.Name, stu))
				}
			}
		case *ast.FuncDecl:
			if node.Name.Name != "" && unicode.IsUpper([]rune(node.Name.Name)[0]) {
				if name := getName(node); name != "" {
					source.Funcs = append(source.Funcs, name)
				}
			}
		case *ast.File:
			source.Package = node.Name.Name
		}

		return true
	})

	return source, nil
}
