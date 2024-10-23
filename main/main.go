package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/vyevs/ansi"
	"github.com/vyevs/trie"
)

func main() {
	if err := myMain(); err != nil {
		fmt.Printf("uh oh, something went wrong: %v", err)
	}
}

func myMain() error {
	dict, err := getDictTimed()
	if err != nil {
		return fmt.Errorf("failed to read dictionary: %v", err)
	}

	trie := buildTrieTimed(dict)

	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a prefix of at least 3 characters: ")

		if haveMore := sc.Scan(); !haveMore {
			break
		}

		line := sc.Text()
		if len(line) < 3 {
			continue
		}

		words := trie.GetStringsWithPrefix(line)

		printResults(words, line)

	}

	if err := sc.Err(); err != nil {
		return fmt.Errorf("scanner error: %v", err)
	}

	return nil
}

func printResults(words []string, prefix string) {
	fmt.Printf("found %d words beginning with "+ansi.FGGreen+prefix+ansi.Clear+"\n", len(words))

	for _, w := range words {
		fmt.Print(ansi.FGGreen + prefix + ansi.FGRed + w[len(prefix):] + ansi.Clear + "\n")
	}
}

func getDictTimed() ([]string, error) {
	defer timeIt(time.Now(), "reading dict")
	return ReadDictionaryFromFile("words_alpha.txt")
}

func buildTrieTimed(dict []string) *trie.Node {
	defer timeIt(time.Now(), "building trie")
	return trie.Build(dict)
}

// ReadDictionaryFromFile uses ReadDictionary to read from the specified file.
func ReadDictionaryFromFile(file string) ([]string, error) {
	bs, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	return ReadDictionary(bytes.NewReader(bs))
}

// ReadDictionary reads a newline-delimited sequence of strings from r and returns them in a slice.
func ReadDictionary(r io.Reader) ([]string, error) {
	sc := bufio.NewScanner(r)

	dict := make([]string, 0, 1<<19)
	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)
		if line != "" {
			dict = append(dict, line)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %v", err)
	}

	return dict, nil
}

func timeIt(start time.Time, s string) {
	fmt.Printf("%s took %v\n", s, time.Since(start))
}
