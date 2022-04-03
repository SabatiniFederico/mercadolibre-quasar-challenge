package service

import (
	"errors"
	"strings"
)

func GetMessage(messages ...[]string) (msg string, err error) {

	shiftedMessages := removeNoiseOfMessages(messages...)

	solution := []string{}

	for i, word := range shiftedMessages[0] {
		word, err := mergeTwoWords(word, messages[1][i])

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

func removeNoiseOfMessages(messages ...[]string) [][]string {
	minLen := minLength(messages...)

	for i, message := range messages {
		messages[i] = message[len(message)-minLen:]
	}
	return messages
}

func minLength(messages ...[]string) (min int) {
	min = len(messages[0])
	for _, message := range messages {
		if min > len(message) {
			min = len(message)
		}
	}
	return min
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
