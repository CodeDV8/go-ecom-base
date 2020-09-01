package tests_test

import (
	"testing"

	EComBase "github.com/codedv8/go-ecom-base"
)

func TestNodeBalance(t *testing.T) {
	first := &EComBase.LinkedTreeNode{
		Key: "K",
	}
	first.Add("O", nil)
	first.Add("C", nil)
	first.Add("R", nil)
	first.Add("S", nil)
	first.Add("Z", nil)
	first.Add("Å", nil)
	first.Add("I", nil)
	first.Add("J", nil)
	first.Add("B", nil)
	first.Add("N", nil)
	first.Add("A", nil)
	first.Add("L", nil)
	first.Add("Y", nil)
	first.Add("Ö", nil)
	first.Add("M", nil)
	first.Add("P", nil)
	first.Add("D", nil)
	first.Add("E", nil)
	first.Add("F", nil)
	first.Add("G", nil)
	first.Add("U", nil)
	first.Add("V", nil)
	first.Add("W", nil)
	first.Add("H", nil)
	first.Add("Q", nil)
	first.Add("Ä", nil)
	first.Add("T", nil)
	first.Add("X", nil)

	balanced, balancedErr, _ := first.Balance(28)
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
