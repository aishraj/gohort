Gohort
==================

Gohort  is a simple URL shortner written in Go.

Its design is based out the [Stack Overflow question](https://stackoverflow.com/questions/742013/how-to-code-a-url-shortener) about writing a URL shortner. It uses [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) for routing requests. 

It provides a RESTful API to create and retrive short URL and their corresponding expanded forms.

Example
===================
In order to create a new short URL:

```curl -X POST http://localhost:8080/api/v1/?base=www.google.com```

In order to retrive the original URL from the shortend URL:

```curl http://localhost:8080/api/v1/\?alias=8CQ```

