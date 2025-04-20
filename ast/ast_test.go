package ast

import (
	"testing"

	"github.com/hirokihello/monkey-interpreter/token"
)

func TestString(t *testing.T) {
	program := &Program{}
	program.Statements = []Statement{
		&LetStatement{
			Token: token.Token{Type: token.LET, Literal: "let"},
			Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "myVar"}, Value: "myVar"},
			Value: &Identifier{Token: token.Token{Type: token.IDENT, Literal: "anotherVar"}, Value: "anotherVar"},
		},
	}

	expected := "let myVar = anotherVar;"
	if program.String() != expected {
		t.Errorf("program.String() = %q, want %q", program.String(), expected)
	}
}
