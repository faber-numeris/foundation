// Package envutils provides helpers for manipulating the process environment
// in tests. Its primary use case is isolating a test from the ambient
// environment: snapshot and clear it with Push, then restore it with Pop.
package envutils

import (
	"errors"
	"os"
	"strings"
	"sync"
)

// ErrEmptyStack is returned by Pop when there is no snapshot to restore, i.e.
// Pop was called more times than Push.
var ErrEmptyStack = errors.New("envutils: Pop called with an empty stack")

var (
	mu    sync.Mutex
	stack [][]string
)

// Push snapshots the current process environment onto an internal stack and
// then clears every environment variable, giving the caller a clean slate.
//
// Each Push must be paired with a matching Pop (deferring Pop right after Push
// is the idiomatic usage) so the original environment is restored.
func Push() error {
	mu.Lock()
	defer mu.Unlock()

	stack = append(stack, os.Environ())
	os.Clearenv()

	return nil
}

// Pop restores the environment captured by the most recent Push, discarding any
// changes made since. It first clears the current environment so variables that
// were set after the matching Push do not leak past it.
//
// It returns ErrEmptyStack if there is no snapshot to restore.
func Pop() error {
	mu.Lock()
	defer mu.Unlock()

	if len(stack) == 0 {
		return ErrEmptyStack
	}

	snapshot := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	os.Clearenv()
	for _, entry := range snapshot {
		key, value, found := strings.Cut(entry, "=")
		if !found {
			continue
		}
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return nil
}
