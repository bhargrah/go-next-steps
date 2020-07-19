package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const FILEPATH = `/Users/rahulbhargava/Desktop/git_repo/go/go-next-steps/02-language-basics/e3/proverbs.txt`


func main() {
	bs, err := ioutil.ReadFile(FILEPATH)
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	proverbs := string(bs) // conversion not casting

	lines := strings.Split(proverbs, "\n")
	for _, l := range lines {
		fmt.Printf("%s\n", l)
		for k, v := range charCount(l) {
			fmt.Printf("'%s'=%d, ", k, v)
		}
		fmt.Print("\n\n")
	}
}

func charCount(line string) map[string]int {
	m := make(map[string]int, 0)
	for _, char := range line {
		m[string(char)] = m[string(char)] + 1
	}
	return m
}
