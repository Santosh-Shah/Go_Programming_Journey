package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter sentence: ")
	sentence, _ := reader.ReadString('\n')

	sentence = reverserWordsInSentence(sentence)
	fmt.Println("Your reversed sentence is: ", sentence)
}

func reverserWordsInSentence(sentence string) string {
	words := strings.Fields(sentence)

	i := 0
	j := len(words) - 1
	for i < j {
		temp := words[i]
		words[i] = words[j]
		words[j] = temp
		i++
		j--
	}

	return strings.Join(words, " ")
}
