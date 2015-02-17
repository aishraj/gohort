package khukuri

import (
	"log"
	"strconv"
)

const Base10 = 10
const Base64 = 64

func CheckSum(id uint64) uint64 {
	numStr := strconv.FormatUint(id, Base10)
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
	return checkSum % Base10
}

func IsCheckSumValid(id uint64) bool {
	return CheckSum(id) == 0
}

func CalculateCheckDigit(partialId uint64) uint64 {
	checkDigit := CheckSum(partialId * Base10)
	if checkDigit == 0 {
		return checkDigit
	}
	return Base10 - checkDigit
}

func digitSum(num uint64) uint64 {
	retNum := uint64(0)
	for num > 0 {
		r := num % Base10
		retNum += r
		num = num / Base10
	}
	return retNum
}

func ValidateAlias(alias string) bool {
	if len(alias) < 2 || alias[0] < '0' || alias[0] > '9' {
		return false
	}
	checkDigit := string(alias[0])
	checkDigitNum, _ := strconv.ParseUint(checkDigit, Base10, Base64)
	actualAlias := string(alias[1:])
	aliasId, ok := DecodeFromBase(actualAlias)
	if ok != nil {
		log.Printf("Unable to decode the value %s . Error is %s", actualAlias, ok.Error())
		return false
	}
	return CalculateCheckDigit(aliasId) == checkDigitNum
}
