package http

import (
	"encoding/xml"
	"github.com/sfomuseum/go-http-opensearch"
	gohttp "net/http"
)

type OpenSearchHandlerOptions struct {
	Description *opensearch.OpenSearchDescription
}

func OpenSearchHandler(opts *OpenSearchHandlerOptions) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		body, err := xml.Marshal(opts.Description)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-Type", "application/opensearchdescription+xml")
		rsp.Write(body)
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
