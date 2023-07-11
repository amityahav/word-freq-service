package internal

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewStore(cfg StoreConfig) *Store {
	s := Store{
		logger:           logrus.New(),
		frequencies:      map[string]*Element{},
		medianStore:      newMedianStore(),
		topKStore:        newTopKStore(cfg.K),
		least:            NewMinHeap(),
		insertionChannel: make(chan []string, cfg.Capacity), //TODO: cfg param for buffering?
	}

	s.logger.SetFormatter(&logrus.JSONFormatter{})

	return &s
}

// Maintain is responsible for keeping the Store up-to-date with the continuously incoming insertion requests.
// it is doing that by behaving like an event loop - fetching insertion events and executing them in a sequential manner.
// NOTE: executing is done sequentially in order to ensure the correctness of the Store's internal data structures
func (s *Store) Maintain() {
	s.logger.
		WithField("component", "store").
		Info("Maintain loop started...")

	for {
		words := <-s.insertionChannel

		s.logger.WithField("component", "store").
			Infof("inserting %v", words)

		s.insertWords(words)
	}
}

// GetStats returns top-5 frequent words, the least frequent word, median frequent word
func (s *Store) GetStats(ctx context.Context) (*Stats, error) {
	s.RLock()
	defer s.RUnlock()

	output := make(chan Stats, 1)

	go s.getStats(output)

	select {
	case <-ctx.Done():
		return nil, errors.Wrap(ctx.Err(), "failed fetching statistics")
	case res := <-output:
		return &res, nil
	}
}

// InsertWords simply adds new words to the insertion channel
func (s *Store) InsertWords(words []string) {
	s.insertionChannel <- words
}

// insertWords inserts/ updates word's frequencies and update Store's internal data structures accordingly
func (s *Store) insertWords(words []string) {
	// locking is done only for synchronizing between writing/reading from store as writing performed sequentially
	s.Lock()
	defer s.Unlock()

	for _, w := range words {
		// inserting to frequencies map
		e := s.insertFreq(w)

		// inserting to top-k heap
		s.topKStore.insert(e)

		// inserting to the median store

		// inserting to the least heap

	}
}

func (s *Store) insertFreq(word string) *Element {
	var e *Element

	if elem, ok := s.frequencies[word]; !ok {
		e = &Element{
			Word:      word,
			Frequency: 1,
		}

		s.frequencies[word] = e
	} else {
		elem.Frequency++
		e = elem
	}

	return e
}

func (s *Store) getStats(output chan Stats) {
	// could be parallelized but those are O(1) operations
	topK, least, median := s.topKStore.getTopK(), s.getLeast(), s.medianStore.getMedian()

	output <- Stats{
		TopK:   topK,
		Least:  least,
		Median: median,
	}
}

func (s *Store) getLeast() Element {
	least := s.least.Peek().(*Element)
	return elementCopy(least)
}
