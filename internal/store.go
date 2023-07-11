package internal

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
	"wordStore/internal/utils"
)

func NewStore(cfg StoreConfig) *Store {
	s := Store{
		logger:           logrus.New(),
		frequencies:      map[string]*utils.Element{},
		medianStore:      newMedianStore(),
		topKStore:        newTopKStore(cfg.K),
		insertionChannel: make(chan []string, cfg.Capacity),
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

// GetStats returns top-k frequent words, the least frequent word and the median frequent word
func (s *Store) GetStats(ctx context.Context) (*Stats, error) {
	s.RLock() // blocking if an insertion is being processed
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

// Insert simply adds new words to the insertion channel
func (s *Store) Insert(words []string) {
	s.insertionChannel <- words
}

// insertWords inserts/updates words frequencies and update Store's internal data structures accordingly
func (s *Store) insertWords(words []string) {
	// locking is done only for synchronizing between writing/reading from store as writing performed sequentially
	s.Lock()
	defer s.Unlock()

	var wg sync.WaitGroup

	elements := map[string]*utils.Element{}

	for _, w := range words {
		// inserting to frequencies map
		elements[w] = s.insertFreq(w)
	}

	// after inserting/updating all new words we can update store's internal data structures concurrently
	wg.Add(1)
	go s.topKStore.insert(elements, &wg)

	wg.Add(1)
	go s.medianStore.insert(elements, &wg)

	wg.Wait()
}

func (s *Store) insertFreq(word string) *utils.Element {
	var e *utils.Element

	if elem, ok := s.frequencies[word]; !ok {
		e = &utils.Element{
			Word:       word,
			Frequency:  1,
			SmallerIdx: -1,
			LargerIdx:  -1,
		}

		s.frequencies[word] = e
	} else {
		elem.Frequency++
		e = elem
	}

	return e
}

func (s *Store) getStats(output chan Stats) {
	topKFreq, leastFreq, medianFreq := s.topKStore.getTopK(), s.medianStore.getLeast(), s.medianStore.getMedian()

	output <- Stats{
		K:      s.topKStore.k,
		TopK:   topKFreq,
		Least:  leastFreq,
		Median: medianFreq,
	}
}
