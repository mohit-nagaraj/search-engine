package utils

import (
	"strings"
	"unicode"
)

// tokenize returns a slice of tokens for the given text.
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		// Split on any character that is not a letter or a number.
		/*calls the FieldsFunc function from the strings package. The FieldsFunc function splits the text string into fields using a custom function. The custom function is defined as an anonymous function that takes a rune as input and returns a boolean value. The rune represents a Unicode code point in Go. The function determines whether a character should be considered a field separator or not.
		This line inside the anonymous function checks if the given rune is not a letter (unicode.IsLetter(r)) and not a number (unicode.IsNumber(r)). If both conditions are true, it returns true, indicating that the character should be considered a field separator.
		This closing parenthesis ends the anonymous function passed as an argument to FieldsFunc.
		This closing brace ends the tokenize function.
		The tokenize function essentially uses the FieldsFunc function from the strings package to split the text string into individual tokens based on the provided custom function. The custom function determines whether a character should be considered a field separator or not based on whether it is a letter or a number. The resulting tokens are returned as a slice of strings.
		*/
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		/*This line inside the anonymous function checks if the given rune is not a letter (unicode.IsLetter(r)) and not a number (unicode.IsNumber(r)). If both conditions are true, it returns true, indicating that the character should be considered a field separator.*/
	})
}

// analyze analyzes the text and returns a slice of tokens.
func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}
