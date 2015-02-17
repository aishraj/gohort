package khukuri

import (
	"errors"
	"strconv"
	"strings"
)

const alphaBetSet = "23456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ-_"
const base = uint64(len(alphaBetSet))

func CheckSum(id uint64) uint64 {
	numStr := strconv.FormatUint(id, 10)
	runes := []rune(numStr)
	checkSum := uint64(0)
	oddSum := uint64(0)
	evenSum := uint64(0)
	for i := len(runes) - 1; i >= 0; i-- {
		item := runes[i]
		digitNum := uint64(item - '0')
		ii := i - len(runes) + 1
		if ii%2 == 0 {
			evenSum += digitNum
		} else {
			sumOfDigits := digitSum(2 * digitNum)
			oddSum += sumOfDigits
		}
	}
	checkSum = evenSum + oddSum
	return checkSum % 10
}

func IsCheckSumValid(id uint64) bool {
	return CheckSum(id) == 0
}

func CalculateCheckDigit(partialId uint64) uint64 {
	checkDigit := CheckSum(partialId * 10)
	if checkDigit == 0 {
		return checkDigit
	}
	return 10 - checkDigit
}

func digitSum(num uint64) uint64 {
	retNum := uint64(0)
	for num > 0 {
		r := num % 10
		retNum += r
		num = num / 10
	}
	return retNum
}

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
