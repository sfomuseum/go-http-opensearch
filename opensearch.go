package opensearch

type OpenSearchImage struct {
	Height int
	Width  int
	URL    string
}

type OpenSearchURL struct {
	Type       string
	Method     string
	Template   string
	Parameters []OpenSearchQueryParameter
}

type OpenSearchQueryParameter struct {
	Name  string
	Value string
}

type OpenSearchDescription struct {
	ShortName   string
	Description string
	Image       OpenSearchImage
	URL         OpenSearchURL
	SearchForm  string
}
