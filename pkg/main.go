package main

import (
	"fmt"
	. "github.com/noodleslove/string_tokenizer/pkg/str_tokenizer"
)

func main() {
	s := "So, it was the night of october 17th. pi was still 3.14. sigh! 2."

	st := NewStrTokenizer()
	st.SetString(s)
	for st.More() {
		t := st.Tokenize()
		fmt.Printf("%8s: ", t.TypeStr())
		t.Display()
		fmt.Printf("\n\n")
	}
}
