package nodes

import (
	"github.com/yuin/goldmark/ast"
)

type ComponentNode struct {
	ast.BaseBlock
	Name string
	Data map[string]any
}

var ComponentNodeKind = ast.NewNodeKind("Component")

func (n *ComponentNode) Kind() ast.NodeKind {
	return ComponentNodeKind
}

func (n *ComponentNode) Dump(b []byte, level int) {
	println("not implemented")
}
