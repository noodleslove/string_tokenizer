package main

import (
	. "fake.com/string_tokenizer/pkg/tokenizer"
	"fmt"
)

func main() {
	s := "So, it was the night of october 17th. pi was still 3.14. sigh! 2."

	st := NewTokenizer()
	st.SetString(s)
	for st.More() {
		t := st.Tokenize()
		fmt.Printf("%8s: ", t.TypeStr())
		t.Display()
		fmt.Printf("\n\n")
	}
}
