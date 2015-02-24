package shortener

import (
	"errors"
	"fmt"
	"github.com/fzzy/radix/redis"
	"log"
	"strconv"
	"time"
)

func SetupRedisConnection(serverHost string, serverDb string, serverport string, timeOutSeconds int) (*redis.Client, error) {
	c, err := redis.DialTimeout("tcp",
		serverHost+":"+serverport, time.Duration(timeOutSeconds)*time.Second)
	performErrorCheck(err)
	c.Cmd("SELECT", serverDb)
	return c, err
}

func performErrorCheck(err error) {
	if err != nil {
		log.Fatal("Error while setting a redis connection. Error is ", err)
	}
}

func LookupAlias(alias string, c *redis.Client) (string, error) {
	defer c.Close()
	if !ValidateAlias(alias) {
		return "", errors.New("Unable to validate the lookup string.")
	}
	alias = alias[1:]
	lookupId, err := DecodeFromBase(alias)
	if err != nil {
		fmt.Println("ERROR!!!!! Can't convert string id")
		return "", err
	}
	s, err := c.Cmd("get", lookupId).Str()
	if err != nil {
		fmt.Println("ERROR!!!!! Can't convert string id")
		return "", err
	}
	return s, nil

}

func StoreUrl(baseUrl string, c *redis.Client) (string, error) {
	defer c.Close()

	res := c.Cmd("incr", "globalCounter")
	performErrorCheck(res.Err)

	currentCounter := res.String()

	idNumber, err := strconv.ParseUint(currentCounter, 10, 64)
	performErrorCheck(err)

	setREsp, err := c.Cmd("setnx", idNumber, baseUrl).Bool()
	performErrorCheck(err)

	if setREsp == false {
		fmt.Println("The ID already exits. ERROR!!")
		//TODO: Need to end the program here. ?
	}

	encodedAlias, ok := EncodeToBase(idNumber)
	if ok != nil {
		return "", ok
	}
	checkDigit := CalculateCheckDigit(idNumber)
	checkDigitStr := strconv.FormatUint(checkDigit, 10)
	return checkDigitStr + encodedAlias, nil
}
