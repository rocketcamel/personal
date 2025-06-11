package renderers

import (
	"bytes"
	"context"
	"fmt"
	"personal/cmd/internal/nodes"
	"personal/templates/components"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type ComponentRenderer struct{}

func (r *ComponentRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(nodes.ComponentNodeKind, r.Render)
}

func (r *ComponentRenderer) Render(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*nodes.ComponentNode)

	if n.Name == "Header" {
		var buf bytes.Buffer
		err := components.Header().Render(context.Background(), &buf)
		if err != nil {
			return ast.WalkStop, err
		}
		_, err = fmt.Fprint(w, buf.String())

		return ast.WalkContinue, err
	}

	return ast.WalkContinue, nil
}
