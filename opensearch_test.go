package opensearch

import (
	_ "fmt"
	"testing"
)

const EXPECTED_MARSHAL string = `<?xml version="1.0" encoding="UTF-8"?>
<OpenSearchDescription xmlns:moz="http://www.mozilla.org/2006/browser/search/" xmlns="http://a9.com/-/spec/opensearch/1.1/"><InputEncoding></InputEncoding><ShortName>Example Search</ShortName><Description>Example Search is an example</Description><Image height="16" width="16">http://localhost:8080/opensearch.jpg</Image><Url type="text/html" method="GET" template="https://localhost:8080/search"><Param name="q" value="{searchTerms}"></Param></Url><moz:searchForm>https://localhost:8080/search</moz:searchForm></OpenSearchDescription>`

func TestOpenSearch(t *testing.T) {

	im := &OpenSearchImage{
		Height: DEFAULT_IMAGE_HEIGHT,
		Width:  DEFAULT_IMAGE_WIDTH,
		URI:    "http://localhost:8080/opensearch.jpg",
	}

	params := []*OpenSearchURLParameter{
		&OpenSearchURLParameter{
			Name:  "q",
			Value: DEFAULT_SEARCHTERMS,
		},
	}

	u := &OpenSearchURL{
		Type:       DEFAULT_URL_TYPE,
		Method:     DEFAULT_URL_METHOD,
		Template:   "https://localhost:8080/search",
		Parameters: params,
	}

	desc := &OpenSearchDescription{
		NSMoz:        NS_MOZ,
		NSOpenSearch: NS_OPENSEARCH,
		ShortName:    "Example Search",
		Description:  "Example Search is an example",
		Image:        im,
		URL:          u,
		SearchForm:   "https://localhost:8080/search",
	}

	enc, err := desc.Marshal()

	if err != nil {
		t.Fatal(err)
	}

	// fmt.Println(string(enc))

	if string(enc) != EXPECTED_MARSHAL {
		t.Fatal("Marshaling failed")
	}
}
