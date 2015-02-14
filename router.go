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
	api.Methods("PUT").MatcherFunc(MultiMatcher).HandlerFunc(ForceCreateHandler)

	fmt.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", r)

}

func RootHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "RootHandler handler")
}

func RedirectToBaseHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Redirect handler")
}

func AliasHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "AliasHandler handler")
}

func BaseHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "BaseHandler handler")
}

func ForceCreateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "ForceCreateHandler handler")
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
