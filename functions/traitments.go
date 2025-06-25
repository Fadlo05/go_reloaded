package functions

func Traitments(s string) []string {
	words := Clean(s)
	tokens := Tokenizer(words)
	tokens = Process(tokens)
	tokens = ChangesWithN(tokens)
	tokens = MergeApostrophes(tokens)
	tokens = Ponctuations(tokens)
	tokens = AToAn(tokens)
	return tokens
}
