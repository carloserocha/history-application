package histories

import (
	"regexp"
	"sort"
	"strings"
)

func RemoveDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func SortDescendingAlphabet(s string) string {
	r := regexp.MustCompile(`\W+`)
	OnlyLetters := r.ReplaceAllString(s, "")

	SplitLetters := strings.Split(OnlyLetters, "")

	sort.Strings(SplitLetters)

	return strings.Join(SplitLetters, "")
}

func SortUniqueLetters(s string) string {
	letters := strings.Split(SortDescendingAlphabet(s), "")
	UniqueLetters := RemoveDuplicates(letters)
	return strings.Join(UniqueLetters, "")
}
