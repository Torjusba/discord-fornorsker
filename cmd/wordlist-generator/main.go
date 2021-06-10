package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/torjusba/discord-fornorsker/pkg/wordlist"
)

func main() {

	var outputPath string
	flag.StringVar(&outputPath, "o", "wordlist.tsv", "Specify the output file")

	fmt.Printf("Generating wordlist at: %s", outputPath)
	var outputFile *os.File
	var err error
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(fmt.Sprintf("Could not create file: \n%s", err))
	}
	defer outputFile.Close()

	wordlist.Init(wordlist.MODE_SIMPLE)
	exportedWordlist := wordlist.Dump()
	outputFile.WriteString(exportedWordlist)
}
