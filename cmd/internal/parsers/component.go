package parsers

import (
	"bytes"
	"personal/cmd/internal/nodes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type ComponentParser struct{}

func (p *ComponentParser) Trigger() []byte {
	return []byte{'@'}
}

func (p *ComponentParser) Open(parent ast.Node, block text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, segment := block.PeekLine()

	if bytes.HasPrefix(line, []byte("@components.Header")) {
		block.Advance(segment.Len())
		return &nodes.ComponentNode{Name: "Header"}, parser.NoChildren
	}

	return nil, parser.NoChildren
}

func (p *ComponentParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	return parser.Close
}

func (p *ComponentParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {}
func (p *ComponentParser) CanInterruptParagraph() bool                                { return false }
func (p *ComponentParser) CanAcceptIndentedLine() bool                                { return true }
