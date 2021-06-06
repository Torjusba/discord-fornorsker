package wordlist

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/torjusba/discord-fornorsker/pkg/logging"
)

var (
	// g_words contains [word]replacement where word is lower case after init
	g_words   map[string]string
	g_inited  bool = false
	g_counter      = 0
)

func Init() {
	g_words = make(map[string]string)

	response, err := http.Get("https://www.sprakradet.no/sprakhjelp/Skriverad/Avloeysarord/")
	if err != nil {
		logging.Error("Error getting word list", err)
		return
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromResponse(response)

	if err != nil {
		logging.Error("Error generating document from word list", err)
		return
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

func GetNeededReplacements(message string) map[string]string {
	relevantReplacements := make(map[string]string)
	lowerCaseMessage := strings.ToLower(message)
	for word, replacement := range g_words {
		if strings.Contains(lowerCaseMessage, word) {
			relevantReplacements[word] = replacement
		}
	}
	return relevantReplacements
}
