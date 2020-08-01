package histories

import "testing"

func TestSortUniqueLetters(t *testing.T) {
    want := "abdglnoprsuxy"
    if got := SortUniqueLetters("sandbox playground"); got != want {
        t.Errorf("SortUniqueLetters() = %q, want %q", got, want)
    }
}

func TestSortDescendingAlphabet(t *testing.T) {
	want := "aabddglnnooprsuxy"
    if got := SortDescendingAlphabet("sandbox playground"); got != want {
        t.Errorf("SortDescendingAlphabet() = %q, want %q", got, want)
    }
}