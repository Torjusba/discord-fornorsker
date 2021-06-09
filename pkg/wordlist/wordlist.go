package wordlist

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Mode string

const (
	MODE_UNDEFINED Mode = "undefined"
	MODE_SIMPLE         = "simple"
)

var (
	// g_words contains [word]replacement where word is lower case after init
	g_words   map[string]string
	g_inited  bool = false
	g_counter      = 0
	g_mode    Mode = MODE_UNDEFINED
)

func Init(mode Mode) {
	g_mode = mode
	defer func() { g_inited = true }()
	g_words = make(map[string]string)

	switch g_mode {
	case MODE_SIMPLE:
		initWordlistFromWebsite()
		return
	default:
		panic(fmt.Sprintf("Mode '%s' not recognized", mode))
	}

}

func initWordlistFromWebsite() {
	response, err := http.Get("https://www.sprakradet.no/sprakhjelp/Skriverad/Avloeysarord/")
	if err != nil {
		panic(fmt.Sprintf("Error getting word list:\n %s", err))
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromResponse(response)

	if err != nil {
		panic(fmt.Sprintf("Error generating document from word list:\n %s", err))
	}

	// Kinda hacky, but this page should only contain one table, which is the
	// one we are interested in
	document.Find("table").First().Find("tr").Each(func(i int, s *goquery.Selection) {
		word := s.Find("td").First().Text()
		replacement := s.Find("td").Last().Text()

		// Skip "letter titles"
		if len([]rune(word)) <= 3 {
			return
		}

		g_words[strings.ToLower(word)] = replacement
	})
}

func GetBorrowedWords() []string {
	if !g_inited {
		panic("Wordlist used before initialization")
	}
	borrowedWords := make([]string, 0)

	switch g_mode {
	case MODE_SIMPLE:
		for borrowedWord := range g_words {
			borrowedWords = append(borrowedWords, borrowedWord)
		}
	}
	return borrowedWords
}

func GetReplacement(word string) string {
	if !g_inited {
		panic("Wordlist used before initialization")
	}
	switch g_mode {
	case MODE_SIMPLE:
		return getReplacement_basic(word)
	default:
		panic(fmt.Sprintf("Replacements in mode '%s' not implemented", g_mode))
	}
}

func getReplacement_basic(word string) string {
	replacement, ok := g_words[strings.ToLower(word)]
	if ok {
		return replacement
	} else {
		panic(fmt.Sprintf("Could not find replacement for word '%s'", word))
	}
}
