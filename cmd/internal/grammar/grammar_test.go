package grammar

import (
	"log"
	"testing"
)

func TestGrammar(t *testing.T) {
	output, err := ComponentParser.ParseString("", "@components.Image(size: 24, width: 24)")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	println(output)
}
