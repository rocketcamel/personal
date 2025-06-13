package grammar

import (
	"testing"
)

func TestGrammar(t *testing.T) {
	output, err := ComponentParser.ParseString("f", "@components.Image(size: 24, width: 24)")
	if err != nil {
		println(err)
	}
	println(output)
}
