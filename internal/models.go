package internal

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	ListenAddress string `yaml:"listen_address"`
}

type medianStore struct {
	smaller *Heap
	larger  *Heap
}

type Store struct {
	sync.RWMutex
	logger *logrus.Logger

	frequencies      map[string]*Element
	ms               *medianStore
	topK             *Heap
	least            *Heap
	insertionChannel chan []string
}

type Stats struct {
	Top5   []Element `json:"top5"`
	Least  Element   `json:"least"`
	Median Element   `json:"median"`
}

type Element struct {
	Word       string `json:"word"`
	Frequency  uint32 `json:"frequency"`
	InSomeHeap bool
}
