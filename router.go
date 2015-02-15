package khukuri

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type UrlMsg struct {
	Url          string `json:shorturl`
	ErrorMessage string `json:errorMessage`
}

func RegisterAndStart() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	shortner := r.Path("/{alias}").Subrouter()
	shortner.Methods("GET").HandlerFunc(RedirectToBaseHandler)

	api := r.PathPrefix("/api/v1/").Subrouter()
	api.Methods("GET").MatcherFunc(AliasMatcher).HandlerFunc(AliasHandler)
	api.Methods("POST").MatcherFunc(BaseMatcher).HandlerFunc(BaseHandler)

	fmt.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", r)

}

func RootHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "RootHandler handler")
}

func RedirectToBaseHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	baseUrl, ok := LookupAlias(vars["alias"])
	if ok != nil {
		fmt.Println("Error while redirecting")
		fmt.Println(ok)
	}
	if baseUrl != "" {
		http.Redirect(rw, r, baseUrl, http.StatusMovedPermanently)
	} else {
		http.NotFound(rw, r)
	}

}

func AliasHandler(rw http.ResponseWriter, r *http.Request) {
	baseUrl, ok := ExtractBaseUrl(r)
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

func ExtractBaseUrl(r *http.Request) (string, error) {
	alias := extractParam(r, "alias")
	return LookupAlias(alias)
}

func BaseHandler(rw http.ResponseWriter, r *http.Request) {
	baseUrl := extractParam(r, "base")
	alias, ok := StoreUrl(baseUrl)
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
