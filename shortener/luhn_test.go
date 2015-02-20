package shortener

import "testing"
import "fmt"

func TestCheckSum(t *testing.T) {
	checkSum := CheckSum(6657949)
	fmt.Println("Checksum is: ", checkSum)

	if checkSum != 5 {
		fmt.Println("The test failed. Checksum is :  ", checkSum)
		t.Fail()
	}
}
