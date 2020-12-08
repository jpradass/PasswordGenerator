package wordlist

// FileJSON is the format of the file for the wordlist
type FileJSON struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}
