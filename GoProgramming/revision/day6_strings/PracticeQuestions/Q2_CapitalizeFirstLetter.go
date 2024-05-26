package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func capitalizeSentence(sentence string) string {
	words := strings.Fields(sentence)
	for i, word := range words {
		words[i] = capitalizedWord(word)
	}

	return strings.Join(words, " ")
}

func capitalizedWord(word string) string {
	if len(word) == 0 {
		return word
	}

	firstLetter := []rune(word)[0]
	return string(unicode.ToUpper(firstLetter)) + word[1:]
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a sentence: ")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	fmt.Println()

	capitalizedSentence := capitalizeSentence(sentence)
	fmt.Println("Capitalized Sentence:", capitalizedSentence)

	//str := "heloo"

	//char := rune(str[0])
	//charValue := string(unicode.ToUpper(char))
	//fmt.Println(charValue)

	//fmt.Println(capitalizeWord("hello"))

	//fmt.Println(str[1:])
	//words := strings.Fields(sentence)
	//fmt.Println(words)
	//fmt.Println(strings.Join(words, " "))

	//count := 0
	//for index, value := range words {
	//	count++
	//	fmt.Println("Index: ", index, ": ", string(value))
	//	//fmt.Println(string(value))
	//}
	//
	//fmt.Println("Count: ", count)

}
