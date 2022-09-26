// Copyright (c) 2022 Marin Atanasov Nikolov <dnaeon@gmail.com>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer
//    in this position and unchanged.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR(S) ``AS IS'' AND ANY EXPRESS OR
// IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
// OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE AUTHOR(S) BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
// THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package deque

import (
        "errors"
        "sync"
)

// ErrEmptyQueue is an error which is returned when attempting to pop
// an item from an empty queue
var ErrEmptyQueue = errors.New("Queue is empty")

type Deque[T any] struct {
        sync.RWMutex
        items []T
}

// New creates a new deque
func New[T any]() *Deque[T] {
        d := &Deque[T]{
                items: make([]T, 0),
        }

        return d
}

// PushBack inserts a new item at the back
func (d *Deque[T]) PushBack(value T) {
        d.Lock()
        defer d.Unlock()
        d.items = append(d.items, value)
}

// PushFront inserts a new item at the front
func (d *Deque[T]) PushFront(value T) {
        d.Lock()
        defer d.Unlock()
        d.items = append([]T{value}, d.items...)
}

// IsEmpty returns true if the deque is empty, false otherwise
func (d *Deque[T]) IsEmpty() bool {
        d.RLock()
        defer d.RUnlock()
        return len(d.items) == 0
}

// PopBack pops an item from the back
func (d *Deque[T]) PopBack() (T, error) {
        var empty T
        if d.IsEmpty() {
                return empty, ErrEmptyQueue
        }

        d.Lock()
        defer d.Unlock()

        size := len(d.items)
        item := d.items[size-1]
        d.items = d.items[:size-1]

        return item, nil
}

// PopFront pops an item from the front
func (d *Deque[T]) PopFront() (T, error) {
        var empty T
        if d.IsEmpty() {
                return empty, ErrEmptyQueue
        }

        d.Lock()
        defer d.Unlock()

        item := d.items[0]
        d.items = d.items[1:]

        return item, nil
}

// Length returns the size of the queue
func (d *Deque[T]) Length() int {
        d.RLock()
        defer d.RUnlock()
        return len(d.items)
}

// PeekFront peeks at the front
func (d *Deque[T]) PeekFront() (T, error) {
        var empty T
        if d.IsEmpty() {
                return empty, ErrEmptyQueue
        }

        d.RLock()
        defer d.RUnlock()

        return d.items[0], nil
}

// PeekBack peeks at the back
func (d *Deque[T]) PeekBack() (T, error) {
        var empty T
        if d.IsEmpty() {
                return empty, ErrEmptyQueue
        }

        d.RLock()
        defer d.RUnlock()

        size := len(d.items)
        return d.items[size-1], nil
}
