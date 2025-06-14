package parsers

import (
	"personal/cmd/internal/grammar"
	"personal/cmd/internal/nodes"
	"reflect"

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
	data := make(map[string]any)

	for _, param := range output.Parameters {
		val := reflect.ValueOf(*param.Value)

		for i := range val.NumField() {
			if !val.Field(i).IsNil() {
				field := val.Field(i).Interface()
				data[param.Key] = field
			}
		}
	}

	block.Advance(segment.Len())
	return &nodes.ComponentNode{Name: output.Name, Data: data}, parser.NoChildren
}

func (p *ComponentParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	return parser.Close
}

func (p *ComponentParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {}
func (p *ComponentParser) CanInterruptParagraph() bool                                { return false }
func (p *ComponentParser) CanAcceptIndentedLine() bool                                { return true }
