package main

import (
	"fmt"
	"io"
	"os"

	"goreloaded/functions"
)

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func main() {
	args := os.Args[1:]
	text, err := readFile(args[0])
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	words := functions.Clean(text)
	tokens := functions.Tokenizer(words)
	tokens = functions.Process(tokens)
	tokens = functions.ChangesWithN(tokens)
	tokens = functions.MergeApostrophes(tokens)
	// result := RebuildText(tokens)
	fmt.Println(tokens)
}
