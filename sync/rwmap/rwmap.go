/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package rwmap

import "sync"

type RWMap struct {
	mu sync.RWMutex
	m  map[string]interface{}
}

func NewRWMap(l int) *RWMap {
	return &RWMap{
		m: make(map[string]interface{}, l),
	}
}

func (m *RWMap) Get(key string) (val interface{}, existed bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, existed = m.m[key]

	return val, existed
}

func (m *RWMap) Set(key string, val interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = val
}

func (m *RWMap) Del(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.m, key)
}

func (m *RWMap) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.m)
}

func (m *RWMap) Each(callback func(key string, val interface{}) bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for k, v := range m.m {
		if !callback(k, v) {
			return
		}
	}
}
