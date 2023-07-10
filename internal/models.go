package internal

import "sync"

type Config struct {
	ListenAddress string `yaml:"listen_address"`
}

type medianStore struct {
	smaller Heap
	larger  Heap
}

type Store struct {
	sync.RWMutex

	ms    medianStore
	topK  Heap
	least Heap
}
