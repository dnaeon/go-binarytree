// Copyright (c) 2022 Marin Atanasov Nikolov <dnaeon@gmail.com>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//  1. Redistributions of source code must retain the above copyright
//     notice, this list of conditions and the following disclaimer
//     in this position and unchanged.
//  2. Redistributions in binary form must reproduce the above copyright
//     notice, this list of conditions and the following disclaimer in the
//     documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR(S) “AS IS” AND ANY EXPRESS OR
// IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
// OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE AUTHOR(S) BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
// THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package binarytree_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/dnaeon/go-binarytree.v1"
)

func TestHeightAndSize(t *testing.T) {
	// Our test tree
	//
	//     __1
	//    /   \
	//   2     3
	//  / \
	// 4   5
	//
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	root.InsertRight(3)
	two.InsertLeft(4)
	five := two.InsertRight(5)

	if root.Size() != 5 {
		t.Fatal("expected tree size should be 5")
	}

	if root.Height() != 2 {
		t.Fatal("expected height from root should be 2")
	}

	if two.Height() != 1 {
		t.Fatal("expected height from node (2) should be 1")
	}

	if five.Height() != 0 {
		t.Fatal("expected height from node (5) should be 0")
	}
}

func TestIsLeafNode(t *testing.T) {
	// Our test tree
	//
	//     __1
	//    /   \
	//   2     3
	//  / \
	// 4   5
	//
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	three := root.InsertRight(3)
	four := two.InsertLeft(4)
	five := two.InsertRight(5)

	if root.IsLeafNode() {
		t.Fatal("root node should not be a leaf")
	}

	if two.IsLeafNode() {
		t.Fatal("node (2) should not be a leaf")
	}

	if !three.IsLeafNode() {
		t.Fatal("node (3) should be a leaf")
	}

	if !four.IsLeafNode() {
		t.Fatal("node (4) should not be a leaf")
	}

	if !five.IsLeafNode() {
		t.Fatal("node (5) should be a leaf")
	}
}

func TestIsFullNode(t *testing.T) {
	// Our test tree
	//
	//     __1
	//    /   \
	//   2     3
	//  / \
	// 4   5
	//
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	three := root.InsertRight(3)
	four := two.InsertLeft(4)
	five := two.InsertRight(5)

	if !root.IsFullNode() {
		t.Fatal("root node should be full")
	}

	if !two.IsFullNode() {
		t.Fatal("node (2) should be full")
	}

	if three.IsFullNode() {
		t.Fatal("node (3) should not be full")
	}

	if four.IsFullNode() {
		t.Fatal("node (4) should not be full")
	}

	if five.IsFullNode() {
		t.Fatal("node (5) should not be full")
	}
}

func TestWalkInOrder(t *testing.T) {
	// Our test tree
	//
	//     __1
	//    /   \
	//   2     3
	//  / \
	// 4   5
	//
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	root.InsertRight(3)
	two.InsertLeft(4)
	two.InsertRight(5)

	result := make([]int, 0)
	wantResult := []int{4, 2, 5, 1, 3}
	walkFunc := func(node *binarytree.Node[int]) error {
		result = append(result, node.Value)
		return nil
	}

	if err := root.WalkInOrder(walkFunc); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, wantResult) {
		t.Fatalf("want in-order values %v, got %v", wantResult, result)
	}
}

func TestWalkPreOrder(t *testing.T) {
	// Our test tree
	//
	//     __1
	//    /   \
	//   2     3
	//  / \
	// 4   5
	//
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	root.InsertRight(3)
	two.InsertLeft(4)
	two.InsertRight(5)

	result := make([]int, 0)
	wantResult := []int{1, 2, 4, 5, 3}
	walkFunc := func(node *binarytree.Node[int]) error {
		result = append(result, node.Value)
		return nil
	}

	if err := root.WalkPreOrder(walkFunc); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantResult, result) {
		t.Fatalf("want pre-order values %v, got %v", wantResult, result)
	}
}

func TestWalkPostOrder(t *testing.T) {
	// Our test tree
	//
	//     __1
	//    /   \
	//   2     3
	//  / \
	// 4   5
	//
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	root.InsertRight(3)
	two.InsertLeft(4)
	two.InsertRight(5)

	result := make([]int, 0)
	wantResult := []int{4, 5, 2, 3, 1}
	walkFunc := func(node *binarytree.Node[int]) error {
		result = append(result, node.Value)
		return nil
	}

	if err := root.WalkPostOrder(walkFunc); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantResult, result) {
		t.Fatalf("want post-order values %v, got %v", wantResult, result)
	}
}

func TestWalkLevelOrder(t *testing.T) {
	// Our test tree
	//
	//     __1
	//    /   \
	//   2     3
	//  / \
	// 4   5
	//
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	root.InsertRight(3)
	two.InsertLeft(4)
	two.InsertRight(5)

	result := make([]int, 0)
	wantResult := []int{1, 2, 3, 4, 5}
	walkFunc := func(node *binarytree.Node[int]) error {
		result = append(result, node.Value)
		return nil
	}

	if err := root.WalkLevelOrder(walkFunc); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantResult, result) {
		t.Fatalf("want level-order values %v, got %v", wantResult, result)
	}
}

func TestSkipNodeHandlers(t *testing.T) {
	// Construct the following simple binary tree
	//
	//     __1
	//    /   \
	//   2     3
	//  / \
	// 4   5
	//
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	root.InsertRight(3)
	two.InsertLeft(4)
	two.InsertRight(5)

	skipFunc := func(n *binarytree.Node[int]) bool {
		// Skip the sub-tree at node (2)
		if n.Value == 2 {
			return true
		}

		return false
	}

	values := make([]int, 0)
	walkFunc := func(n *binarytree.Node[int]) error {
		values = append(values, n.Value)
		return nil
	}

	root.AddSkipNodeFunc(skipFunc)
	if err := root.WalkInOrder(walkFunc); err != nil {
		t.Fatal(err)
	}

	wantValues := []int{1, 3}
	if !reflect.DeepEqual(values, wantValues) {
		t.Fatalf("want in-order values %v, got %v", wantValues, values)
	}
}

func TestFindNode(t *testing.T) {
	// Construct the following simple binary tree
	//
	//     __1
	//    /   \
	//   2     3
	//  / \
	// 4   5
	//
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	root.InsertRight(3)
	two.InsertLeft(4)
	two.InsertRight(5)

	goodPredicate := func(n *binarytree.Node[int]) bool {
		if n.Value == 2 {
			return true
		}
		return false
	}

	node, ok := root.FindNode(goodPredicate)
	if !ok {
		t.Fatal("unable to find node (2)")
	}

	// The node we are looking for should have left and right
	// children
	if node.Left == nil || node.Right == nil {
		t.Fatal("node (2) does not have left or right children")
	}

	// No node will match is supposed to match with this predicate
	badPredicate := func(n *binarytree.Node[int]) bool {
		return false
	}

	if _, ok := root.FindNode(badPredicate); ok {
		t.Fatal("no node is supposed to match the predicate")
	}
}

func TestIsFullTree(t *testing.T) {
	// Our test tree
	//
	//     __1
	//    /   \
	//   2     3
	//  / \
	// 4   5
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	root.InsertRight(3)
	two.InsertLeft(4)
	two.InsertRight(5)

	if !root.IsFullTree() {
		t.Fatal("tree should be full")
	}
}

func TestIsNotFullTree(t *testing.T) {
	// Our test tree
	//
	//     __1
	//    /
	//   2
	//  / \
	// 4   5
	//
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	two.InsertLeft(4)
	two.InsertRight(5)

	if root.IsFullTree() {
		t.Fatal("tree should not be full")
	}
}

func TestTreeIsDegenerateTree(t *testing.T) {
	// Our test tree
	//
	//     1
	//    /
	//   2
	//    \
	//     3
	//    /
	//   4
	//
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	three := two.InsertRight(3)
	three.InsertLeft(4)

	if !root.IsDegenerateTree() {
		t.Fatal("tree should be degenerate")
	}
}

func TestTreeIsNotDegenerate(t *testing.T) {
	// Our test tree
	//
	//     __1
	//    /
	//   2
	//  / \
	// 4   5
	//
	root := binarytree.NewNode(1)
	two := root.InsertLeft(2)
	two.InsertLeft(4)
	two.InsertRight(5)

	if root.IsDegenerateTree() {
		t.Fatal("tree should not be degenerate")
	}
}

func TestIsBalancedTree(t *testing.T) {
	// Unbalanced tree
	//
	//     1
	//    /
	//   2
	//  /
	// 3
	//
	unbalanced_root := binarytree.NewNode(1)
	two := unbalanced_root.InsertLeft(2)
	two.InsertLeft(3)

	if unbalanced_root.IsBalancedTree() {
		t.Fatal("tree should not be balanced")
	}

	// Another unbalanced tree
	//
	//         __1
	//        /   \
	//     __2     3
	//    /   \
	//   4     5
	//  / \
	// 6   7
	unbalanced_root = binarytree.NewNode(1)
	unbalanced_root.InsertRight(3)
	two = unbalanced_root.InsertLeft(2)
	two.InsertRight(5)
	four := two.InsertLeft(4)
	four.InsertLeft(6)
	four.InsertRight(7)

	if unbalanced_root.IsBalancedTree() {
		t.Fatal("tree should not be balanced")
	}

	// A single root node is a balanced tree
	leaf := binarytree.NewNode(1)
	if !leaf.IsBalancedTree() {
		t.Fatal("single leaf node should be balanced")
	}

	// Yet another unbalanced tree
	//
	//     __1
	//    /
	//   2
	//  / \
	// 4   5
	//
	unbalanced_root = binarytree.NewNode(1)
	two = unbalanced_root.InsertLeft(2)
	two.InsertLeft(4)
	two.InsertRight(5)

	if unbalanced_root.IsBalancedTree() {
		t.Fatal("tree should not be balanced")
	}

	// The sub-tree with root node (2) is balanced
	if !two.IsBalancedTree() {
		t.Fatal("tree with root (2) should be balanced")
	}

	// A balanced tree
	//
	//   1__
	//  /   \
	// 2     3
	//      / \
	//     4   5
	balanced_root := binarytree.NewNode(1)
	balanced_root.InsertLeft(2)
	three := balanced_root.InsertRight(3)
	three.InsertLeft(4)
	three.InsertRight(5)

	if !balanced_root.IsBalancedTree() {
		t.Fatal("tree should be balanced")
	}
}

func TestIsCompleteTree(t *testing.T) {
	// A complete binary tree
	//
	//    1
	//   / \
	//  2   3
	//
	root := binarytree.NewNode(1)
	root.InsertLeft(2)
	root.InsertRight(3)

	if !root.IsCompleteTree() {
		t.Fatal("tree should be complete")
	}

	// A complete binary tree
	//
	//      1
	//     / \
	//    2   3
	//   /
	//  4
	//
	root = binarytree.NewNode(1)
	root.InsertRight(3)
	two := root.InsertLeft(2)
	two.InsertLeft(4)

	if !root.IsCompleteTree() {
		t.Fatal("tree should be complete")
	}

	// A complete binary tree
	//
	//     __1__
	//    /     \
	//   2       3
	//  / \     /
	// 4   5   6
	//
	root = binarytree.NewNode(1)
	two = root.InsertLeft(2)
	two.InsertLeft(4)
	two.InsertRight(5)
	three := root.InsertRight(3)
	three.InsertLeft(6)

	if !root.IsCompleteTree() {
		t.Fatal("tree should be complete")
	}

	// Not complete binary tree
	//
	//     __1_
	//    /    \
	//   2      3
	//  / \      \
	// 4   5      6
	//
	root = binarytree.NewNode(1)
	two = root.InsertLeft(2)
	two.InsertLeft(4)
	two.InsertRight(5)
	three = root.InsertRight(3)
	three.InsertRight(6)

	if root.IsCompleteTree() {
		t.Fatal("tree should not be complete")
	}

	// Not complete binary tree
	//
	//     1
	//    /
	//   2
	//  /
	// 3
	//
	root = binarytree.NewNode(1)
	two = root.InsertLeft(2)
	two.InsertLeft(3)

	if root.IsCompleteTree() {
		t.Fatal("tree should not be complete")
	}

	// Not complete binary tree
	//
	//   __1__
	//  /     \
	// 2       3
	//  \     / \
	//   4   5   6
	root = binarytree.NewNode(1)
	two = root.InsertLeft(2)
	two.InsertRight(4)
	three = root.InsertRight(3)
	three.InsertLeft(5)
	three.InsertRight(6)

	if root.IsCompleteTree() {
		t.Fatal("tree should not be complete")
	}

	// Not complete binary tree
	//
	//   1__
	//  /   \
	// 2     3
	//      / \
	//     4   5
	root = binarytree.NewNode(1)
	root.InsertLeft(2)
	three = root.InsertRight(3)
	three.InsertLeft(4)
	three.InsertRight(5)

	if root.IsCompleteTree() {
		t.Fatal("tree should not be complete")
	}

	// A complete binary tree with a single root node
	root = binarytree.NewNode(1)
	if !root.IsCompleteTree() {
		t.Fatal("tree should be complete")
	}
}

func TestIsPerfectTree(t *testing.T) {
	// A perfect binary tree
	//
	//    1
	//   / \
	//  2   3
	//
	root := binarytree.NewNode(1)
	root.InsertLeft(2)
	root.InsertRight(3)

	if !root.IsPerfectTree() {
		t.Fatal("tree should be perfect")
	}

	// A non-perfect binary tree
	//
	//      1
	//     / \
	//    2   3
	//   /
	//  4
	//
	root = binarytree.NewNode(1)
	root.InsertRight(3)
	two := root.InsertLeft(2)
	two.InsertLeft(4)

	if root.IsPerfectTree() {
		t.Fatal("tree should not be perfect")
	}

	// A non-perfect binary tree
	//
	//     __1__
	//    /     \
	//   2       3
	//  / \     /
	// 4   5   6
	//
	root = binarytree.NewNode(1)
	two = root.InsertLeft(2)
	two.InsertLeft(4)
	two.InsertRight(5)
	three := root.InsertRight(3)
	three.InsertLeft(6)

	if root.IsPerfectTree() {
		t.Fatal("tree should not be perfect")
	}

	// A perfect binary tree
	//
	//     __1__
	//    /     \
	//   2       3
	//  / \     / \
	// 4   5   6   7
	//
	root = binarytree.NewNode(1)
	two = root.InsertLeft(2)
	two.InsertLeft(4)
	two.InsertRight(5)
	three = root.InsertRight(3)
	three.InsertLeft(6)
	three.InsertRight(7)

	if !root.IsPerfectTree() {
		t.Fatal("tree should be perfect")
	}

	// A non-perfect binary tree
	//
	//     1
	//    /
	//   2
	//  /
	// 3
	//
	root = binarytree.NewNode(1)
	two = root.InsertLeft(2)
	two.InsertLeft(3)

	if root.IsPerfectTree() {
		t.Fatal("tree should not be perfect")
	}

	// A non-perfect binary tree
	//
	//   1__
	//  /   \
	// 2     3
	//      / \
	//     4   5
	root = binarytree.NewNode(1)
	root.InsertLeft(2)
	three = root.InsertRight(3)
	three.InsertLeft(4)
	three.InsertRight(5)

	if root.IsPerfectTree() {
		t.Fatal("tree should not be perfect")
	}

	// A perfect binary tree with a single root node
	root = binarytree.NewNode(1)
	if !root.IsPerfectTree() {
		t.Fatal("tree should be perfect")
	}
}

func TestIsBinarySearchTree(t *testing.T) {
	// A valid BST
	//
	//    2
	//   / \
	//  1   3
	//
	root := binarytree.NewNode(2)
	root.InsertLeft(1)
	root.InsertRight(3)

	if !root.IsBinarySearchTree(binarytree.IntComparator) {
		t.Fatal("tree should be BST")
	}

	// Invalid BST
	//
	//    1
	//   / \
	//  2   3
	//
	root = binarytree.NewNode(1)
	root.InsertLeft(2)
	root.InsertRight(3)

	if root.IsBinarySearchTree(binarytree.IntComparator) {
		t.Fatal("tree should not be BST")
	}

	// Invalid BST
	//
	//      1
	//     / \
	//    2   3
	//   /
	//  4
	//
	root = binarytree.NewNode(1)
	root.InsertRight(3)
	two := root.InsertLeft(2)
	two.InsertLeft(4)

	if root.IsBinarySearchTree(binarytree.IntComparator) {
		t.Fatal("tree should not be BST")
	}

	// A valid BST
	//
	//     ______8
	//    /       \
	//   3__       10___
	//  /   \           \
	// 1     6          _14
	//      / \        /
	//     4   7      13
	//
	root = binarytree.NewNode(8)
	three := root.InsertLeft(3)
	three.InsertLeft(1)
	six := three.InsertRight(6)
	six.InsertLeft(4)
	six.InsertRight(7)
	ten := root.InsertRight(10)
	fourteen := ten.InsertRight(14)
	fourteen.InsertLeft(13)

	if !root.IsBinarySearchTree(binarytree.IntComparator) {
		t.Fatal("tree should be BST")
	}

	// A tree with a single root node is a valid BST
	root = binarytree.NewNode(1)
	if !root.IsBinarySearchTree(binarytree.IntComparator) {
		t.Fatal("tree should be BST")
	}

	// A valid BST
	//
	//   B
	//  / \
	// A   C
	//
	str_root := binarytree.NewNode("B")
	str_root.InsertLeft("A")
	str_root.InsertRight("C")

	if !str_root.IsBinarySearchTree(binarytree.StringComparator) {
		t.Fatal("tree should be BST")
	}

	// Invalid BST
	//
	//     A
	//    / \
	//   B   C
	//  /
	// D
	str_root = binarytree.NewNode("A")
	str_root.InsertRight("C")
	b := str_root.InsertLeft("B")
	b.InsertLeft("D")

	if str_root.IsBinarySearchTree(binarytree.StringComparator) {
		t.Fatal("tree should not be BST")
	}
}

func TestNodeAttributes(t *testing.T) {
	root := binarytree.NewNode(1)

	if root.GetDotAttributes() != "" {
		t.Fatal("node is expected to have no attributes")
	}

	root.AddAttribute("color", "green")

	wantAttrs := "color=green"
	gotAttrs := root.GetDotAttributes()
	if gotAttrs != wantAttrs {
		t.Fatal("node attributes mismatch")
	}
}

func TestWriteDot(t *testing.T) {
	// Our test tree
	//
	//   1__
	//  /   \
	// 2     3
	//      / \
	//     4   5
	root := binarytree.NewNode(1)
	root.InsertLeft(2)
	three := root.InsertRight(3)
	three.InsertLeft(4)
	three.InsertRight(5)

	var buf bytes.Buffer
	if err := root.WriteDot(&buf); err != nil {
		t.Fatal(err)
	}

	output := buf.String()
	if output == "" {
		t.Fatal("got empty dot representation")
	}

	if !strings.HasPrefix(output, "digraph {") {
		t.Fatal("missing dot prefix")
	}

	if !strings.HasSuffix(output, "}\n") {
		t.Fatal("missing dot suffix")
	}
}
