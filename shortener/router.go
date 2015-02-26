package shortener

import (
	"encoding/json"
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UrlMsg struct {
	Url          string `json:shorturl`
	ErrorMessage string `json:errorMessage`
}

var hostRedis string = ""
var dbRedis string = ""
var portRedis string = ""
var timeOutRedis int = 10

func RegisterAndStart(redisHost string, redisDatabase string, redisPort string, serverPort string, timeOutSeconds int) {
	hostRedis = redisHost
	dbRedis = redisDatabase
	portRedis = redisPort
	timeOutRedis = timeOutSeconds
	serverPort = ":" + serverPort

	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	shortener := r.Path("/{alias}").Subrouter()
	shortener.Methods("GET").HandlerFunc(RedirectToBaseHandler)

	api := r.PathPrefix("/api/v1/").Subrouter()
	api.Methods("GET").MatcherFunc(AliasMatcher).HandlerFunc(AliasHandler)
	api.Methods("POST").MatcherFunc(BaseMatcher).HandlerFunc(BaseHandler)

	log.Println("Server starting on port ", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, r))

}

func RootHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println(r.UserAgent(), "  ", r.Method)
	fmt.Fprint(rw, "Welcome to the shortener URL shortener v0.01")
}

func RedirectToBaseHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(r.UserAgent(), "  ", r.Method, r.URL)
	c, err := SetupRedisConnection(hostRedis, dbRedis, portRedis, timeOutRedis)
	if err != nil {
		log.Fatal("Unable to setup a redis connection", err)
	}

	baseUrl, ok := LookupAlias(vars["alias"], c)
	if ok != nil {
		log.Println("Error while redirecting")
		log.Println(ok)
		http.NotFound(rw, r)
	}
	if baseUrl != "" {
		http.Redirect(rw, r, baseUrl, http.StatusMovedPermanently)
	} else {
		http.NotFound(rw, r)
	}

}

func AliasHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println(r.UserAgent(), "  ", r.Method, r.URL)
	c, err := SetupRedisConnection(hostRedis, dbRedis, portRedis, timeOutRedis)
	if err != nil {
		log.Fatal("Unable to setup a redis connection", err)
	}

	baseUrl, ok := ExtractBaseUrl(r, c)
	urlMessage := UrlMsg{"", ""}
	if ok == nil {
		urlMessage = UrlMsg{baseUrl, ""}
	} else {
		urlMessage = UrlMsg{baseUrl, ok.Error()}
	}
	js, _ := json.Marshal(urlMessage)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func ExtractBaseUrl(r *http.Request, c *redis.Client) (string, error) {
	alias := extractParam(r, "alias")
	return LookupAlias(alias, c)
}

func BaseHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println(r.UserAgent(), "  ", r.Method, r.URL)
	c, err := SetupRedisConnection(hostRedis, dbRedis, portRedis, timeOutRedis)
	if err != nil {
		log.Fatal("Unable to setup a redis connection", err)
	}

	baseUrl := extractParam(r, "base")
	alias, ok := StoreUrl(baseUrl, c)
	if ok != nil {
		http.Error(rw, ok.Error(), http.StatusInternalServerError)
	} else {
		urlMessage := UrlMsg{alias, ""}
		js, _ := json.Marshal(urlMessage)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(js)
	}
}

func AliasMatcher(r *http.Request, rm *mux.RouteMatch) bool {
	queryParams := r.URL.Query()
	return queryParams.Get("alias") != ""
}

func BaseMatcher(r *http.Request, rm *mux.RouteMatch) bool {
	queryParams := r.URL.Query()
	return queryParams.Get("base") != ""
}

func MultiMatcher(r *http.Request, rm *mux.RouteMatch) bool {
	return AliasMatcher(r, rm) && BaseMatcher(r, rm)
}

func extractParam(r *http.Request, a string) string {
	queryPrams := r.URL.Query()
	return queryPrams.Get(a)
}
