package internal

import (
	"personal/cmd/internal/parsers"
	"personal/cmd/internal/renderers"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type MDComponent struct{}

func NewComponentParser() parser.BlockParser {
	return &parsers.ComponentParser{}
}

func NewComponentRenderer() renderer.NodeRenderer {
	return &renderers.ComponentRenderer{}
}

func (e *MDComponent) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithBlockParsers(
		util.Prioritized(NewComponentParser(), 100),
	))

	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewComponentRenderer(), 100),
	))
}
