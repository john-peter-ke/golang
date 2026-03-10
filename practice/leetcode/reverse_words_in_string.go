package leetcode

import "strings"

func reverseWords(s string) string {
	var sentence string
	var word string
	for i := 0; i < len(s); i++ {
		val := string(s[i])
		if strings.Contains(val, " ") {
			if len(word) > 0 {
				if len(sentence) > 0 {
					sentence = word + " " + sentence
				} else {
					sentence = word
				}
				word = ""
			}
			continue
		}

		word += val
	}

	if len(word) > 0 {
		if len(sentence) > 0 {
			sentence = word + " " + sentence
		} else {
			sentence = word
		}
		word = ""
	}

	return sentence
}
