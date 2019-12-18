package main

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-http-opensearch"
	"github.com/sfomuseum/go-http-opensearch/http"
	"html/template"
	"log"
	gohttp "net/http"
	"path/filepath"
)

var index_html = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>go-http-opensearch</title>
  </head>
  <body>This is the index page</body>
</html>
`

var search_html = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>go-http-opensearch</title>
  </head>
  <body>This is the search page. Query term is <q>{{ .Term }}</q>.</body>
</html>
`

func IndexHandler() (gohttp.Handler, error) {

	t := template.New("index")
	t, err := t.Parse(index_html)

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		err := t.Execute(rsp, nil)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}

type SearchVars struct {
	Term string
}

func SearchHandler() (gohttp.Handler, error) {

	t := template.New("search")
	t, err := t.Parse(search_html)

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		q := req.URL.Query()
		term := q.Get("term")

		vars := SearchVars{
			Term: term,
		}

		err := t.Execute(rsp, vars)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}

func main() {

	host := flag.String("host", "localhost", "...")
	port := flag.Int("port", 8080, "...")

	flag.Parse()

	path_index := "/"
	path_search := "/search/"
	path_opensearch := "/opensearch/"

	endpoint := fmt.Sprintf("%s:%d", *host, *port)

	searchform_url := filepath.Join(endpoint, path_search)

	desc_opts := &opensearch.BasicDescriptionOptions{
		Name:     "Example Search",
		Description:   "Example Search is an example",		
		QueryParameter: "term",
		ImageURI:    "http://localhost:8080/opensearch.jpg",
		SearchTemplate: fmt.Sprintf("http://%s", searchform_url),
		SearchForm: fmt.Sprintf("http://%s", searchform_url),
	}

	desc, err := opensearch.BasicDescription(desc_opts)

	if err != nil {
		log.Fatal(err)
	}
	
	index_handler, err := IndexHandler()

	if err != nil {
		log.Fatal(err)
	}

	search_handler, err := SearchHandler()

	if err != nil {
		log.Fatal(err)
	}

	opensearch_opts := &http.OpenSearchHandlerOptions{
		Description: desc,
	}

	opensearch_handler, err := http.OpenSearchHandler(opensearch_opts)

	if err != nil {
		log.Fatal(err)
	}

	plugins := map[string]*opensearch.OpenSearchDescription{
		path_opensearch: desc,
	}

	plugins_opts := &http.AppendPluginsOptions{
		Plugins: plugins,
	}

	index_handler = http.AppendPluginsHandler(index_handler, plugins_opts)
	search_handler = http.AppendPluginsHandler(search_handler, plugins_opts)

	mux := gohttp.NewServeMux()

	mux.Handle(path_index, index_handler)
	mux.Handle(path_search, search_handler)
	mux.Handle(path_opensearch, opensearch_handler)

	log.Printf("Listening for requests on %s\n", endpoint)

	err = gohttp.ListenAndServe(endpoint, mux)

	if err != nil {
		log.Fatal(err)
	}
}
