/*
1. Read the file
2. Split the file into lines
3. Strip out punctuation, double spaces, and make everything lowercase
4. Split the lines into words
5. Count the words
6. Sort the words by count
7. Print the top 10 words
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// 1. Read the file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 2. Split the file into lines
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// 3. Strip out punctuation, double spaces, and make everything lowercase
	re1 := regexp.MustCompile(`[[:punct:]]`)
	re2 := regexp.MustCompile(`\s+`)
	for i, line := range lines {
		lines[i] = re1.ReplaceAllString(line, "")
		lines[i] = re2.ReplaceAllString(line, " ")
		lines[i] = strings.ToLower(line)
	}

	// 4. Split the lines into words
	words := make([]string, 0)
	for _, line := range lines {
		words = append(words, strings.Split(line, " ")...)
	}

	// 5. Count the words
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}

	// 6. Sort the words by count
	type wordCount struct {
		word  string
		count int
	}

	wordCountsSlice := make([]wordCount, 0)
	for word, count := range wordCounts {
		wordCountsSlice = append(wordCountsSlice, wordCount{word, count})
	}

	sort.Slice(wordCountsSlice, func(i, j int) bool {
		return wordCountsSlice[i].count > wordCountsSlice[j].count
	})

	// 7. Print the top 10 words
	for i := 0; i < 10; i++ {
		fmt.Println(wordCountsSlice[i].word, wordCountsSlice[i].count)
	}
}
