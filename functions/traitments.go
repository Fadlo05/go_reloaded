package functions

func Traitments(s string) []string {
	words := Clean(s)
	tokens := Tokenizer(words)
	tokens = Process(tokens)
	resStr := Ponctuations(tokens)
	tokens = MergeApostrophes(resStr)
	tokens = AToAn(tokens)
	return tokens
}
