package khukuri

import "testing"
import "fmt"

func TestConnectorSet(t *testing.T) {
	//Okay need a better way to do this.
	// Shouldn't be connecting to redis.
	a, ok := StoreUrl("http://www.amazon.com")
	if ok != nil {
		fmt.Println(ok)
		t.Fail()
	}
	fmt.Println("Short URL iD is:", a)
}

func TestConnectorGet(t *testing.T) {
	//Okay need a better way to do this.
	// Shouldn't be connecting to redis.
	a, ok := LookupAlias("33")
	if ok != nil {
		fmt.Println(ok)
		t.Fail()
	}
	fmt.Println("Long URL is:", a)
}
