package functions

func Traitments(s string) []string {
	words := Clean(s)
	tokens := Tokenizer(words)
	tokens = Process(tokens)
	tokens = Ponctuations(tokens)
	tokens = ChangesWithN(tokens)
	tokens = MergeApostrophes(tokens)
	tokens = AToAn(tokens)
	return tokens
}
