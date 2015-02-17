package khukuri

import "testing"
import "fmt"

func TestCheckSum(t *testing.T) {
	checkSum := CheckSum(999999999999)
	fmt.Println("Checksum is: ", checkSum)

	if checkSum != 8 {
		fmt.Println("The test failed. Checksum is :  ", checkSum)
		t.Fail()
	}
}
