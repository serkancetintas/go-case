package store

import (
	"encoding/json"
	"go-case/internal"
	"go-case/internal/domain"
	"go-case/pkg"
	"sync"
	"time"
)

type Store struct {
	data domain.Data
	sync.RWMutex
}

func NewStore() *Store {
	store := &Store{data: map[string]string{}}
	pkg.CreateDirectory(internal.StoreFilePath)
	pkg.ReadJsonFromFile((*map[string]string)(&store.data), internal.StoreFilePath)

	store.save()

	return store
}

func (store *Store) Get(key string) string {
	store.RLock()
	defer store.RUnlock()
	return store.data[key]
}

func (store *Store) Set(key string, value string) {
	store.Lock()
	defer store.Unlock()
	store.data[key] = value
}

func (store *Store) Flush() {
	store.Lock()
	defer store.Unlock()
	for k := range store.data {
		delete(store.data, k)
	}
	pkg.DeleteFilesFromDir(internal.StoreFilePath)
}

func (store *Store) load() {
	store.Lock()
	defer store.Unlock()
	jsonStr, _ := json.Marshal(store.data)

	pkg.CreateFile(jsonStr,internal.StoreFilePath)
}

func (store *Store) save() chan struct{} {
	ticker := time.NewTicker(time.Duration(internal.IntervalTime) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				store.load()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}