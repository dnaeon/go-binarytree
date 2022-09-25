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

package binarytree

import (
        deque "gopkg.in/dnaeon/go-deque.v1"
)

// WalkFunc is the type of the function which will be invoked while
// visiting a node from the binary tree.
type WalkFunc[T any] func(node *Node[T]) error

// SkipNodeFunc is a function which returns true, if the currently
// being visited node should be skipped.
type SkipNodeFunc[T any] func(node *Node[T]) bool

// Node represents a node from a binary tree
type Node[T any] struct {
        // Value is the value of the node
        Value T
        // Left child of the node
        Left *Node[T]
        // Right child of the node
        Right *Node[T]

        // A list of function handlers, which specify whether a node
        // should be skipped or not during tree walking.
        skipNodeFuncs []SkipNodeFunc[T]
}

// NewNode creates a new node
func NewNode[T any](value T) *Node[T] {
        node := &Node[T]{
                Value: value,
                Left:  nil,
                Right: nil,
        }

        return node
}

// InsertLeft inserts a new node to the left
func (n *Node[T]) InsertLeft(value T) *Node[T] {
        left := NewNode(value)
        n.Left = left

        return left
}

// InsertRight inserts a new node to the right
func (n *Node[T]) InsertRight(value T) *Node[T] {
        right := NewNode(value)
        n.Right = right

        return right
}

// WalkInOrder performs an iterative In-order walking of the binary
// tree - Left-Node-Right (LNR)
func (n *Node[T]) WalkInOrder(walkFunc WalkFunc[T]) error {
        stack := deque.New[*Node[T]]()
        node := n

        for node != nil || !stack.IsEmpty() {
                for node != nil {
                        if n.shouldSkipNode(node) {
                                node = nil
                                break
                        }
                        stack.PushFront(node)
                        node = node.Left
                }

                if !stack.IsEmpty() {
                        item, err := stack.PopFront()
                        if err != nil {
                                panic(err)
                        }

                        if err := walkFunc(item); err != nil {
                                return err
                        }

                        node = item.Right
                }
        }

        return nil
}

// WalkPreOrder performs an iterative Pre-order walking of the
// binary tree - Node-Left-Right (NLR)
func (n *Node[T]) WalkPreOrder(walkFunc WalkFunc[T]) error {
        stack := deque.New[*Node[T]]()
        stack.PushFront(n)

        for !stack.IsEmpty() {
                node, err := stack.PopFront()
                if err != nil {
                        panic(err)
                }

                if n.shouldSkipNode(node) {
                        continue
                }

                if err := walkFunc(node); err != nil {
                        return err
                }

                if node.Right != nil {
                        stack.PushFront(node.Right)
                }

                if node.Left != nil {
                        stack.PushFront(node.Left)
                }
        }

        return nil
}

// WalkPostOrder performs an iterative Post-order walking of the
// binary tree - Left-Right-Node (LRN)
func (n *Node[T]) WalkPostOrder(walkFunc WalkFunc[T]) error {
        stack := deque.New[*Node[T]]()
        result := deque.New[*Node[T]]()
        stack.PushFront(n)

        for !stack.IsEmpty() {
                node, err := stack.PopFront()
                if err != nil {
                        panic(err)
                }

                if n.shouldSkipNode(node) {
                        continue
                }

                if node.Left != nil {
                        stack.PushFront(node.Left)
                }
                if node.Right != nil {
                        stack.PushFront(node.Right)
                }

                result.PushFront(node)
        }

        for !result.IsEmpty() {
                node, err := result.PopFront()
                if err != nil {
                        return err
                }
                if err := walkFunc(node); err != nil {
                        return err
                }
        }

        return nil
}

// Size returns the size of the tree
func (n *Node[T]) Size() int {
        size := 0
        walkFunc := func(n *Node[T]) error {
                size++
                return nil
        }
        n.WalkInOrder(walkFunc)

        return size
}

type nodeHeight[T any] struct {
        node   *Node[T]
        height int
}

// Height returns the height of the tree
func (n *Node[T]) Height() int {
        max_height := 0
        root := &nodeHeight[T]{
                node:   n,
                height: 0,
        }
        stack := deque.New[*nodeHeight[T]]()
        stack.PushFront(root)

        for !stack.IsEmpty() {
                item, err := stack.PopFront()
                if err != nil {
                        panic(err)
                }

                if item.height > max_height {
                        max_height = item.height
                }

                if item.node.Right != nil {
                        right := &nodeHeight[T]{
                                node:   item.node.Right,
                                height: item.height + 1,
                        }
                        stack.PushFront(right)
                }
                if item.node.Left != nil {
                        left := &nodeHeight[T]{
                                node:   item.node.Left,
                                height: item.height + 1,
                        }
                        stack.PushFront(left)
                }
        }

        return max_height
}

// IsLeaf returns true, if the node is a leaf, false otherwise.
func (n *Node[T]) IsLeaf() bool {
        return n.Left == nil && n.Right == nil
}

// AddSkipNodeFunc adds a new handler for determining whether a
// node from the tree should be skipped while traversing it.
func (n *Node[T]) AddSkipNodeFunc(handler SkipNodeFunc[T]) {
        n.skipNodeFuncs = append(n.skipNodeFuncs, handler)
}

// shouldSkipNode applies the list of SkipNodeFunc handlers in
// order to determine whether a node should be skipped while walking
// the tree.
func (n *Node[T]) shouldSkipNode(node *Node[T]) bool {
        for _, handler := range n.skipNodeFuncs {
                if handler(node) {
                        return true
                }
        }

        return false
}
