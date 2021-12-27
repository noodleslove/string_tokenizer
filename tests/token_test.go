package tests

import (
	. "fake.com/string_tokenizer/tokenizer"
	"testing"
)

func TestToken(t *testing.T) {
	tok := NewToken("a", 1)
	if tok == nil || tok.Type() != 1 || tok.TokenStr() != "a" {
		t.Errorf("NewToken incorrect")
	}
}

func TestTypeStr(t *testing.T) {
	alpha := NewToken("a", 1)
	num := NewToken("1", 2)
	punc := NewToken("!", 3)
	space := NewToken(" ", 4)
	unno := NewToken("愛", -1)
	if alpha.TypeStr() != "ALPHA" || num.TypeStr() != "NUMBER" || punc.TypeStr() != "PUNC" || space.TypeStr() != "SPACE" || unno.TypeStr() != "UNKNOWN" {
		t.Errorf("TypeStr incorrect")
	}
}
