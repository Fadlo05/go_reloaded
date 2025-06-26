package main

import (
	"fmt"
	"io"
	"os"
	"strings"

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
	if strings.HasSuffix(args[0], ".go") || strings.HasSuffix(args[1], ".go"){
		fmt.Println("Invalid Input.")
		return
	}
	text, err := readFile(args[0])
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	str := functions.Traitments(text)
	result := functions.RebuildText(str)
	outputFile := args[1]
	err = os.WriteFile(outputFile, []byte(result), 0o644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}
	outputData, err := os.ReadFile(outputFile)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier de sortie: %v\n", err)
		return
	}
	fmt.Println(string(outputData))
}
