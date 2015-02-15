package khukuri

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

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
	//fmt.Fprint(rw, "Redirect handler")
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
	fmt.Fprint(rw, "AliasHandler handler ")
	baseUrl := ExtractBaseUrl(r)
	fmt.Fprint(rw, baseUrl)

}

func ExtractBaseUrl(r *http.Request) string {
	alias := extractParam(r, "alias")
	baseUrl, ok := LookupAlias(alias)
	if ok != nil {
		fmt.Println("ERROR: ", ok)
	}
	return baseUrl
}

func BaseHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "BaseHandler handler ")
	baseUrl := extractParam(r, "base")
	alias, ok := StoreUrl(baseUrl)
	if ok != nil {
		fmt.Println("ERROR: ", ok)
	}
	fmt.Fprint(rw, alias)
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
