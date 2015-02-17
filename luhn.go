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
