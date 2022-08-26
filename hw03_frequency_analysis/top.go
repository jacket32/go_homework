package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(str string) []string {
	if str == "" {
		return nil
	}

	str = strings.ToLower(str)

	onlyWords := regexp.MustCompile(`[A-Za-zА-ЯЁа-яё][^.:,!" \s]*`)
	s := onlyWords.FindAllString(str, -1)

	wCount := make(map[string]int)

	sort.Strings(s)

	for _, v := range s {
		wCount[v]++
	}

	words := make([]string, 0, len(wCount))
	for word := range wCount {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool {
		if wCount[words[i]] == wCount[words[j]] {
			return words[i] < words[j]
		}

		return wCount[words[i]] > wCount[words[j]]
	})

	length := len(words)
	if length > 10 {
		return words[:10]
	}

	return words[:length]
}
