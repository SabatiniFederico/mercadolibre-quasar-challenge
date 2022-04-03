package service

import (
	"errors"
	"strings"
)

func GetMessage(messages ...[]string) (msg string, err error) {

	len := len(messages[0])
	solution := []string{}

	for i := 0; i < len; i++ {
		word, err := mergeTwoWords(messages[0][i], messages[1][i])

		if err != nil {
			return "", err
		}

		word, err2 := mergeTwoWords(word, messages[2][i])

		if err2 != nil {
			return "", err2
		}

		solution = append(solution, word)
	}

	return strings.Join(solution, " "), nil
}
func mergeTwoWords(word1, word2 string) (string, error) {
	if word1 != "" {
		if word1 != word2 && word2 != "" {
			return word1, errors.New("contradiction found")
		}
		return word1, nil
	}
	return word2, nil
}
