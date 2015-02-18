package gohort 

import (
	"errors"
	"strings"
)

const alphaBetSet = "23456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ-_"
const base = uint64(len(alphaBetSet))

func EncodeToBase(seedNumber uint64) (string, error) {
	encodedString := ""
	if seedNumber <= 0 {
		return "", errors.New("Argument cannot zero or less.")
	}
	for seedNumber > 0 {
		numAtIndex := seedNumber % base
		charAtIndex := string(alphaBetSet[numAtIndex])
		encodedString += charAtIndex
		seedNumber = seedNumber / base
	}
	return reverse(encodedString), nil
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func DecodeFromBase(encodedString string) (uint64, error) {
	if len(encodedString) == 0 {
		return 0, errors.New("Argument cannot empty.")
	}
	decodedVal := uint64(0)
	for _, c := range encodedString {
		decodedVal = (decodedVal * base) + uint64(strings.Index(alphaBetSet, string(c)))
	}
	return decodedVal, nil
}
