package internal

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"wordStore/internal/utils"
)

func TestStore(t *testing.T) {
	var (
		median uint32
		least  uint32
		topK   []*utils.Element
		words  []string
		stats  *Stats
		err    error
	)

	cfg := StoreConfig{
		K:        5,
		Capacity: 1000,
	}

	ctx := context.Background()
	store := NewStore(cfg)

	t.Run("insertion/fetching", func(t *testing.T) {
		checkEqual := func(stats *Stats, store *Store) {
			median = naiveMedian(store.frequencies)
			require.Equal(t, median, stats.Median, "wrong median")

			least = naiveLeast(store.frequencies)
			require.Equal(t, least, stats.Least, "wrong least")

			topK = naiveTopK(store.frequencies, cfg.K)
			require.Equal(t, true, compareElements(topK, stats.TopK), "wrong top-k")
		}

		words = []string{}
		store.insertWords(words)

		stats, err = store.GetStats(ctx)
		require.NoError(t, err)

		checkEqual(stats, store)

		words = []string{"ball", "eggs", "pool", "dart", "ball", "ball"}
		store.insertWords(words)

		stats, err = store.GetStats(ctx)
		require.NoError(t, err)

		checkEqual(stats, store)

		words = []string{"table", "eggs", "pool", "mouse", "ball", "eggs"}
		store.insertWords(words)

		stats, err = store.GetStats(ctx)
		require.NoError(t, err)

		checkEqual(stats, store)

		words = []string{"table", "mouse"}
		store.insertWords(words)

		stats, err = store.GetStats(ctx)
		require.NoError(t, err)

		checkEqual(stats, store)
	})

	t.Run("benchmarking", func(t *testing.T) {
		// TODO
	})
}
