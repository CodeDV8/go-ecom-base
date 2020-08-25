package tests_test

import (
	EComBase "github.com/codedv8/go-ecom-base"
	"testing"
)

func TestBalance(t *testing.T) {
	tree := &EComBase.LinkedTree{}
	tree.Add("O", nil)
	tree.Add("C", nil)
	tree.Add("K", nil)
	tree.Add("R", nil)
	tree.Add("S", nil)
	tree.Add("Z", nil)
	tree.Add("Å", nil)
	tree.Add("I", nil)
	tree.Add("J", nil)
	tree.Add("B", nil)
	tree.Add("N", nil)
	tree.Add("A", nil)
	tree.Add("L", nil)
	tree.Add("Y", nil)
	tree.Add("Ö", nil)
	tree.Add("M", nil)
	tree.Add("P", nil)
	tree.Add("D", nil)
	tree.Add("E", nil)
	tree.Add("F", nil)
	tree.Add("G", nil)
	tree.Add("U", nil)
	tree.Add("V", nil)
	tree.Add("W", nil)
	tree.Add("H", nil)
	tree.Add("Q", nil)
	tree.Add("Ä", nil)
	tree.Add("T", nil)
	tree.Add("X", nil)

	balanced, balancedErr := tree.Balance()
	if balancedErr != nil {
		t.Error("Balance failed")
	} else if balanced == nil {
		t.Error("Balance should not return nil if successful")
	} else {
		t.Logf("The key of the new root is %s", balanced.Key)
		left := balanced.Left
		for left != nil {
			t.Logf("Left key is %s", left.Key)
			left = left.Left
		}
		right := balanced.Right
		for right != nil {
			t.Logf("Right key is %s", right.Key)
			right = right.Right
		}
	}
}
