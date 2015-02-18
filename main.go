package main

import (
	"flag"
	gohort "github.com/aishraj/gohort/shortner"
	"log"
	"runtime"
	"strconv"
)

func main() {

	redisHost := flag.String("rhost", "localhost", "Host on which Redis is running")
	redisPortInt := flag.Int("rport", 6379, "Port on which Redis is running")
	redisTimeOutSeconds := flag.Int("timeout", 10, "Timeout for Redis connection in seconds")
	serverPortInt := flag.Int("sport", 8080, "Port for the HTTP server")

	serverPort := strconv.Itoa(*serverPortInt)
	redisPort := strconv.Itoa(*redisPortInt)
	cpus := flag.Int("cpus", runtime.NumCPU(), "Number of CPUs to use")

	flag.Parse()
	runtime.GOMAXPROCS(*cpus)

	log.Printf("Starting the server with properties. Redis host %s "+
		"Redis port number %s Redis Timeout seconds %d HTTP Server port %s",
		*redisHost, redisPort, *redisTimeOutSeconds, serverPort)
	gohort.RegisterAndStart(*redisHost, redisPort, serverPort, *redisTimeOutSeconds)
}
