package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func findAnagrams(words *[]string) *map[string][]string {
	anagrams := make(map[string][]string)
	wordsSet := make(map[string]struct{})
	wordsSorted := make(map[string]string)

	for _, word := range *words {
		wordLower := strings.ToLower(word)
		if _, ok := wordsSet[wordLower]; ok {
			continue
		}

		wordAsc := sortString(wordLower)
		if wordOrig, ok := wordsSorted[wordAsc]; ok {
			anagrams[wordOrig] = append(anagrams[wordOrig], wordLower)
		} else {
			anagrams[wordLower] = []string{wordLower}
			wordsSorted[wordAsc] = wordLower
		}

		wordsSet[wordLower] = struct{}{}
	}

	for key, value := range anagrams {
		if len(value) < 2 {
			delete(anagrams, key)
		} else {
			sort.Strings(value)
		}
	}

	return &anagrams
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток", "тОк", "кто"}
	anagrams := findAnagrams(&words)
	for key, value := range *anagrams {
		fmt.Printf("%s: %v\n", key, value)
	}
}
