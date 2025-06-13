package parsers

import (
	"personal/cmd/internal/grammar"
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

	output, err := grammar.ComponentParser.ParseBytes("", line)
	if err != nil {
		return nil, parser.NoChildren
	}

	block.Advance(segment.Len())
	return &nodes.ComponentNode{Name: output.Name}, parser.NoChildren
}

func (p *ComponentParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	return parser.Close
}

func (p *ComponentParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {}
func (p *ComponentParser) CanInterruptParagraph() bool                                { return false }
func (p *ComponentParser) CanAcceptIndentedLine() bool                                { return true }
