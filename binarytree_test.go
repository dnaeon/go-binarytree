package binarytree_test

import (
        "reflect"
        "testing"

        "gopkg.in/dnaeon/go-binarytree.v1"
)

func TestBinaryTree(t *testing.T) {
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

        if root.IsLeaf() {
                t.Fatal("root node should not be a leaf")
        }

        if !five.IsLeaf() {
                t.Fatal("node (5) should be a leaf")
        }

        collectorFunc := func(values *[]int) binarytree.WalkFunc[int] {
                walkFunc := func(node *binarytree.Node[int]) error {
                        *values = append(*values, node.Value)
                        return nil
                }

                return walkFunc
        }

        inOrderValues := make([]int, 0)
        preOrderValues := make([]int, 0)
        postOrderValues := make([]int, 0)

        wantInOrderValues := []int{4, 2, 5, 1, 3}
        wantPreOrderValues := []int{1, 2, 4, 5, 3}
        wantPostOrderValues := []int{4, 5, 2, 3, 1}

        if err := root.WalkInOrder(collectorFunc(&inOrderValues)); err != nil {
                t.Fatal(err)
        }

        if err := root.WalkPreOrder(collectorFunc(&preOrderValues)); err != nil {
                t.Fatal(err)
        }

        if err := root.WalkPostOrder(collectorFunc(&postOrderValues)); err != nil {
                t.Fatal(err)
        }

        if !reflect.DeepEqual(inOrderValues, wantInOrderValues) {
                t.Fatalf("want in-order values %v, got %v", wantInOrderValues, inOrderValues)
        }

        if !reflect.DeepEqual(preOrderValues, wantPreOrderValues) {
                t.Fatalf("want pre-order values %v, got %v", wantPreOrderValues, preOrderValues)
        }

        if !reflect.DeepEqual(postOrderValues, wantPostOrderValues) {
                t.Fatalf("want post-order values %v, got %v", wantPostOrderValues, postOrderValues)
        }
}
