package internal

import (
	"github.com/sirupsen/logrus"
	"sync"
	"wordStore/internal/utils"
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
	smaller *utils.Heap
	larger  *utils.Heap
}

type topKStore struct {
	k    int
	heap *utils.Heap
}

type Store struct {
	sync.RWMutex
	logger *logrus.Logger

	frequencies      map[string]*utils.Element
	medianStore      *medianStore
	topKStore        *topKStore
	insertionChannel chan []string
}

type Stats struct {
	K      int              `json:"k"`
	TopK   []*utils.Element `json:"topK,omitempty"`
	Least  uint32           `json:"least,omitempty"`
	Median uint32           `json:"median,omitempty"`
}
