package recursivemutex

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // Goroutine ID of the lock currently held
	recursion int32 // Amount goroutine reentries with lock
}

func NewRecursiveMutex() *RecursiveMutex {
	return &RecursiveMutex{}
}

type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64 // Token of the lock currently held
	recursion int32 // Amount goroutine reentries with lock
}

func NewTokenRecursiveMutex(token int64) *TokenRecursiveMutex {
	return &TokenRecursiveMutex{token: token}
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++

		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	m.recursion--
	if m.recursion != 0 {
		return
	}
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

func (t *TokenRecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&t.token) == token {
		t.recursion++

		return
	}
	t.Mutex.Lock()
	atomic.StoreInt64(&t.token, token)
	t.recursion = 1
}

func (t *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&t.token) != token {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", t.token, token))
	}
	t.recursion--
	if t.recursion != 0 {
		return
	}
	atomic.StoreInt64(&t.token, -1)
	t.Mutex.Unlock()
}
