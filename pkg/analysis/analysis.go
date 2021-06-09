package analysis

import (
	"fmt"
	"regexp"
	"strings"
)

// Match 1 whitespace or common punctuation
const WORD_SEPARATOR_REGEX string = "[\\s,\\./()\\?\\\\\\*\\+\\-~#&\\^%]+"

func doesSentenceContainWord(sentence, word string) bool {
	// Pad sentence for regex match on edges
	lowerCaseSentence := " " + strings.ToLower(sentence) + " "
	lowerCaseWord := strings.ToLower(word)

	// Complex check for words in a context
	wordRegex := WORD_SEPARATOR_REGEX + lowerCaseWord + WORD_SEPARATOR_REGEX
	didRegexMatch, err := regexp.Match(wordRegex, []byte(lowerCaseSentence))
	if err != nil {
		panic(fmt.Sprintf("Error matching regex to sentence: %s", err))
	}
	return didRegexMatch
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
