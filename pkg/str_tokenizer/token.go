package str_tokenizer

import (
	"fmt"
)

type Token struct {
	str   string
	_type int
}

func NewToken(str string, _type int) *Token {
	p := Token{str: str, _type: _type}
	return &p
}

func (t *Token) Type() int {
	return t._type
}

func (t *Token) TypeStr() string {
	switch t._type {
	case 1:
		return "ALPHA"
	case 2:
		return "NUMBER"
	case 3:
		return "PUNC"
	case 4:
		return "SPACE"
	default:
		return "UNKNOWN"
	}
}

func (t *Token) TokenStr() string {
	return t.str
}

func (t *Token) Display() {
	fmt.Printf("|%s|", t.str)
}
