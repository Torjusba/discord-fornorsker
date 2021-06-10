package wordlist

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Mode string

const (
	MODE_UNDEFINED Mode = "undefined"
	MODE_SIMPLE         = "simple"
	MODE_FILE           = "file"
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

	switch g_mode {
	case MODE_SIMPLE:
		createWordlistFromWebsite()
		return
	case MODE_FILE:
		g_words = createWordlistFromFile("custom-wordlist.tsv")
		return
	default:
		panic(fmt.Sprintf("Mode '%s' not recognized", mode))
	}
}

func createWordlistFromWebsite() {
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
func createWordlistFromFile(filepath string) map[string]string {

	file, err := os.Open(filepath)
	if err != nil {
		panic(fmt.Sprintf("Could not open file: %s", err))
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("Could not read from file: %s", err))
	}

	newWordlist := make(map[string]string)
	wordlistAsString := string(data)

	lines := strings.Split(wordlistAsString, "\n")
	for _, line := range lines[1:] {
		wordAndReplacement := strings.Split(line, "\t")
		if len(wordAndReplacement) == 2 {
			word := wordAndReplacement[0]
			replacement := wordAndReplacement[1]
			newWordlist[strings.ToLower(word)] = replacement
		}
	}
	return newWordlist
}

func GetBorrowedWords() []string {
	if !g_inited {
		panic("Wordlist used before initialization")
	}
	borrowedWords := make([]string, 0)

	for borrowedWord := range g_words {
		borrowedWords = append(borrowedWords, borrowedWord)
	}
	return borrowedWords
}

func GetReplacement(word string) string {
	if !g_inited {
		panic("Wordlist used before initialization")
	}
	return getReplacement_basic(word)
}

func getReplacement_basic(word string) string {
	replacement, ok := g_words[strings.ToLower(word)]
	if ok {
		return replacement
	} else {
		panic(fmt.Sprintf("Could not find replacement for word '%s'", word))
	}
}

// Bar (|) Separated Values since commas and semicolons are used in the data
func Dump() string {
	var dataToExport string
	for word, replacement := range g_words {
		dataToExport = dataToExport + fmt.Sprintf("\n%s\t%s", word, replacement)
	}
	return dataToExport
}
