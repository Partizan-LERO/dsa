package ransom_note

import "fmt"

// check if you can create word from an array of letters

func solution(letters []string, word string) bool {
	w := map[string]int{}
	for _, letter := range word {
		w[string(letter)] = w[string(letter)] + 1
	}

	for _, letter := range letters {
		res, ok := w[letter]
		if ok {
			w[letter] = res - 1
		}
	}

	for k := range w {
		if w[k] > 0 {
			return false
		}
	}

	return true
}

func solutionImproved(letters []string, word string) bool {
	w := map[string]int{}
	for _, letter := range letters {
		w[letter] = w[letter] + 1
	}

	for _, letter := range word {
		res, ok := w[string(letter)]
		if !ok || res <= 0 {
			return false
		}

		w[string(letter)] = res - 1
	}

	return true
}

func Solution() {
	letters := []string{"a", "b", "c", "d", "e"}
	word := "bed"
	fmt.Println(solutionImproved(letters, word))

	word = "cat"
	fmt.Println(solutionImproved(letters, word))

	word = "bbed"
	fmt.Println(solutionImproved(letters, word))
}
