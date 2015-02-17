package khukuri

import "strconv"

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

func ValidateAlias(alias string) bool {
	if len(alias) < 2 || alias[0] < '0' || alias[0] > '9' {
		return false
	}
	checkDigit := string(alias[0])
	checkDigitNum, _ := strconv.ParseUint(checkDigit, 10, 64)
	actualAlias := string(alias[1:])
	aliasId, ok := DecodeFromBase(actualAlias)
	if ok != nil {
		//TODO log
		return false
	}
	return CalculateCheckDigit(aliasId) == checkDigitNum
}
