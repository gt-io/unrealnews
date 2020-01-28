package main

import (
	"log"
	"sync"
)

var (
	cacheMap map[string]struct{}
	lock     sync.RWMutex
)

// MaxContainerSize max cache size
const MaxContainerSize = 10000

func cacheFlush() error {
	lock.Lock()
	defer lock.Unlock()

	cacheMap = make(map[string]struct{})

	log.Println("flush cache..")
	return nil
}

func cacheRegister(url string) {
	lock.Lock()
	defer lock.Unlock()

	cacheMap[url] = struct{}{}
	log.Println("add cache :", url)
}

func cacheSize() int {
	lock.RLock()
	defer lock.RUnlock()

	return len(cacheMap)
}

func cacheExist(url string) bool {
	lock.RLock()
	defer lock.RUnlock()

	_, ok := cacheMap[url]
	return ok
}
