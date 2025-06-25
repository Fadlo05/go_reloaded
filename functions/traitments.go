package functions

import "fmt"

func Traitments(s string) []string {
	words := Clean(s)
	tokens := Tokenizer(words)
	tokens = Process(tokens)
	tokens = ChangesWithN(tokens)
	fmt.Println(tokens)
	tokens = MergeApostrophes(tokens)
	tokens = Ponctuations(tokens)
	tokens = AToAn(tokens)
	return tokens
}
