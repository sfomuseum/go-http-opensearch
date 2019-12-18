package http

import (
	"errors"
	"github.com/sfomuseum/go-http-opensearch"
	gohttp "net/http"
	"text/template"
)

type OpenSearchHandlerOptions struct {
	Description *opensearch.OpenSearchDescription
	Templates   *template.Template
}

type OpenSearchVars struct {
	Description *opensearch.OpenSearchDescription
}

func OpenSearchHandler(opts *OpenSearchHandlerOptions) (gohttp.Handler, error) {

	t := opts.Templates.Lookup("opensearch")

	if t == nil {
		return nil, errors.New("Missing 'opensearch' template")
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		vars := OpenSearchVars{
			Description: opts.Description,
		}

		rsp.Header().Set("Content-Type", "application/opensearchdescription+xml")

		err := t.Execute(rsp, vars)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
