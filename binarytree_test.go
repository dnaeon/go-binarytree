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

func TestNodeAttributes(t *testing.T) {
	root := binarytree.NewNode(1)

	if root.GetDotAttributes() != "" {
		t.Fatal("node is expected to have no attributes")
	}

	root.AddAttribute("color", "green")
	root.AddAttribute("fillcolor", "green")

	wantAttrs := "color=green fillcolor=green"
	if root.GetDotAttributes() != wantAttrs {
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
