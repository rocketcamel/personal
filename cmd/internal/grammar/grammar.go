package grammar

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Component struct {
	Pos lexer.Position

	At         string      `parser:"@'@'"`
	Namespace  string      `parser:"@Ident '.'"`
	Name       string      `parser:"@Ident"`
	Parameters *ParamsList `parser:"'(' @@? ')'"`
}

type ParamsList struct {
	Params []*Param `parser:"@@ ( ',' @@ )*"`
}

type Param struct {
	Key   string `parser:"@Ident ':'"`
	Value *Value `parser:"@@"`
}

type Value struct {
	String *string  `parser:" @String"`
	Number *float64 `parser:"| @Number"`
	Ident  *string  `parser:"| @Ident"`
}

var ComponentLexer = lexer.MustSimple([]lexer.SimpleRule{
	{Name: "Ident", Pattern: `[a-zA-Z_][a-zA-Z_0-9]*`},
	{Name: "String", Pattern: `"[^"]*"`},
	{Name: "Number", Pattern: `[-+]?[.0-9]+\b`},
	{Name: "Punct", Pattern: `\[|]|[-!()+/*=,.@:]`},
	{Name: "Whitespace", Pattern: `\s+`},
})
var ComponentParser = participle.MustBuild[Component](participle.Lexer(ComponentLexer), participle.Unquote("String"), participle.Elide("Whitespace"))
