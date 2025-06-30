package functions

func Traitments(s string) []string {
	words := Clean(s)
	tokens := Tokenizer(words)
	tokens = AToAn(tokens)
	tokens = Process(tokens)
	tokens = ChangesWithN(tokens)
	resStr := Ponctuations(tokens)
	tokens = MergeApostrophes(resStr)
	return tokens
}
