package internal

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	ListenAddress string      `yaml:"listen_address"`
	Store         StoreConfig `yaml:"store"`
}

type StoreConfig struct {
	K        int `yaml:"k"`
	Capacity int `yaml:"capacity"`
}

type medianStore struct {
	smaller *Heap
	larger  *Heap
}

type topKStore struct {
	k    int
	heap *Heap
}

type leastStore struct {
	// TODO: maybe we can use only the min from the smaller heap of the media store
}

type Store struct {
	sync.RWMutex
	logger *logrus.Logger

	frequencies      map[string]*Element
	medianStore      *medianStore
	topKStore        *topKStore
	least            *Heap
	insertionChannel chan []string
}

type Stats struct {
	TopK   []Element `json:"topK"`
	Least  Element   `json:"least"`
	Median Element   `json:"median"`
}

type Element struct {
	Word      string `json:"word"`
	Frequency uint32 `json:"frequency"`
}
