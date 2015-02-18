package gohort 

import "testing"
import "log"

func TestEnCode(t *testing.T) {
	actual, err := EncodeToBase(6657949)
	if err != nil {
		t.Error("Error while trying to encode 12345. Error is", err)
	}
	if actual != "_cP3" {
		t.Error("Expected _cP3, got", actual)
	}
}

func TestDecodeFail(t *testing.T) {
	_, err := DecodeFromBase("")
	if err == nil {
		t.Error("Expecting error while trying to decode an empty string. Got no error")
	}
}

func TestDecode(t *testing.T) {
	result, err := DecodeFromBase("_cP3")
	if err != nil {
		t.Error("Expecting NO error while trying to decode 5N6. Got ", err)
	}
	if result != 6657949 {
		log.Println(result)
		t.Fail()
	}
}
