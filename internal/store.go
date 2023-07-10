package internal

func NewStore() *Store {
	return &Store{}
}

// GetStats returns top-5 frequent words, the least frequent word, median frequent word
func GetStats() {

}

// Maintain is responsible for keeping the Store up-to-date with the continuously incoming insertion requests.
// it is doing that by behaving like an event loop - fetching insertion events and executing them in a sequential manner.
// NOTE: executing is done sequentially in order to ensure the correctness of the Store's internal data structures
func Maintain() {

}
