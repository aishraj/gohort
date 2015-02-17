Gohort
==================

Gohort  is a simple URL shortner written in Go.

Its design is based out the [Stack Overflow question](https://stackoverflow.com/questions/742013/how-to-code-a-url-shortener) about writing a URL shortner. It uses [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) for routing requests. 

It provides a RESTful API to create and retrive short URL and their corresponding expanded forms.

Running Gohort
=================

Gohort requires a working Redis installation.

Once you have a working Redis installation, go get the project from Github.

```go get github.com/aishraj/gohort```

Now change into the project directory and run 
```go build```

Next run the executable

```./gohort -cpus=1 -rhost="localhost" -rport=6379 -sport=8090 -timeout=10```


Example
===================
In order to create a new short URL:

```curl -X POST http://localhost:8080/api/v1/?base=www.google.com```

In order to retrive the original URL from the shortend URL:

```curl http://localhost:8080/api/v1/\?alias=8CQ```

