package shortener

import "testing"
import "fmt"

func TestConnectorSet(t *testing.T) {
	//Okay need a better way to do this.
	// Shouldn't be connecting to redis.
	c, err := SetupRedisConnection("localhost", "6379", 10)
	if err != nil {
		fmt.Println("Error getting redis connections")
	}
	a, ok := StoreUrl("http://www.amazon.com", c)
	if ok != nil {
		fmt.Println(ok)
		t.Fail()
	}
	fmt.Println("Short URL iD is:", a)
}

func TestConnectorGet(t *testing.T) {
	//Okay need a better way to do this.
	// Shouldn't be connecting to redis.

	c, err := SetupRedisConnection("localhost", "6379", 10)
	if err != nil {
		fmt.Println("Error getting redis connections")
	}
	a, ok := LookupAlias("1c", c)
	if ok != nil {
		fmt.Println(ok)
		t.Fail()
	}
	fmt.Println("Long URL is:", a)
}
