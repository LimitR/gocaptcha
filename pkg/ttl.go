package pkg

import (
	"sync"
	"time"
)

type TTLMap struct {
	sync.RWMutex
	m   map[string]int64
	Ttl int
}

func New(ln int, maxTTL int) (m *TTLMap) {
	m = &TTLMap{m: make(map[string]int64, 10), Ttl: maxTTL}
	go func() {
		for now := range time.Tick(time.Millisecond) {
			m.RLock()
			for k, v := range m.m {
				if now.Unix()-v > int64(maxTTL) {
					delete(m.m, k)
				}
			}
			m.RUnlock()
		}
	}()
	return m
}

func (m *TTLMap) MaxTtl() int {
	return m.Ttl
}

func (m *TTLMap) Put(k string) {
	if m.Ttl == 0 {
		return
	}
	m.RLock()
	_, ok := m.m[k]
	if !ok {
		m.m[k] = time.Now().Unix()
	}
	m.RUnlock()
}

func (m *TTLMap) Get(k string) *int64 {
	if m.Ttl == 0 {
		return nil
	}
	m.RLock()
	if it, ok := m.m[k]; ok {
		return &it
	}
	m.RUnlock()
	return nil

}
