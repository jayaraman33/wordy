package wordy

import (
	"strconv"
	"strings"
)

func Answer(s string) (int, bool) {
	s = strings.TrimSuffix(s, "?")
	var word string
	var err error
	var val int

	arr := strings.Split(s, " ")
	i, curr := 2, 0
	if i >= len(arr) {
		return 0, false
	}

	word = arr[i]

	if val, err = strconv.Atoi(word); err != nil {
		return 0, false
	}
	curr = val

	for {
		i++
		if i == len(arr) {
			return curr, true
		}
		word = arr[i]

		switch word {
		case "plus":
			i++
			if i == len(arr) {
				return 0, false
			}
			word = arr[i]
			if val, err = strconv.Atoi(word); err != nil {
				return 0, false
			}
			curr += val
		case "minus":
			i++
			if i == len(arr) {
				return 0, false
			}
			word = arr[i]
			if val, err = strconv.Atoi(word); err != nil {
				return 0, false
			}
			curr -= val
		case "multiplied":
			i += 2
			if i >= len(arr) {
				return 0, false
			}
			word = arr[i]
			if val, err = strconv.Atoi(word); err != nil {
				return 0, false
			}
			curr *= val
		case "divided":
			i += 2
			if i >= len(arr) {
				return 0, false
			}
			word = arr[i]
			if val, err = strconv.Atoi(word); err != nil {
				return 0, false
			}
			curr /= val
		default:
			return 0, false
		}
	}
}
