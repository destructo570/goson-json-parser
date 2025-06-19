package models

import (
	"fmt"
	"goson/enum"
)

type Token struct {
	Type    enum.TokenType
	Lexeme  string
	Line    int
	Literal any
}

func (t Token) ToString() string {
	return fmt.Sprintf("%s %s %v", t.Type, t.Lexeme, t.Literal)
}
