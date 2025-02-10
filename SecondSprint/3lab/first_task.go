package main

import "sync"

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func (s *SafeMap) Get(key string) interface{} {
	s.mux.Lock()
	data, exists := s.m[key]
	s.mux.Unlock()
	if exists {
		return data
	}
	return nil
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	s.m[key] = value
	s.mux.Unlock()
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]interface{})}
}
