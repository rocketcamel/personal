package renderers

import (
	"bytes"
	"context"
	"fmt"
	"personal/cmd/internal/nodes"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var components_map = map[string]func() templ.Component{}

type ComponentRenderer struct{}

func (r *ComponentRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(nodes.ComponentNodeKind, r.Render)
}

func (r *ComponentRenderer) Render(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*nodes.ComponentNode)

	component := components_map[n.Name]
	if component != nil {
		var buf bytes.Buffer
		err := component().Render(context.Background(), &buf)
		if err != nil {
			return ast.WalkStop, err
		}
		_, err = fmt.Fprint(w, buf.String())

		return ast.WalkContinue, err
	}

	return ast.WalkContinue, nil
}
