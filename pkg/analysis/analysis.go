package analysis

import "strings"

func doesSentenceContainWord(sentence, word string) bool {
	lowerCaseSentence := strings.ToLower(sentence)
	lowerCaseWord := strings.ToLower(word)
	return strings.Contains(lowerCaseSentence, lowerCaseWord)
}

func WhichWordsAreInSentence(sentence string, words []string) []string {
	containedWords := make([]string, 0)
	for _, word := range words {
		if doesSentenceContainWord(sentence, word) {
			containedWords = append(containedWords, word)
		}
	}
	return containedWords
}
