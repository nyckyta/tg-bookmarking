package categorization

type KeyWord struct {
	Text string
}

// KeyWordsFetcher is an interface for fetching keywords from by a certain request
type KeyWordsFetcher interface {
	Fetch(text string) ([]string, error)
}