package tests_test

import (
	EComBase "github.com/codedv8/go-ecom-base"
	"testing"
)

func TestTreeNode(t *testing.T) {
	node := &EComBase.LinkedTreeNode{
		Key:  "F",
		Data: "xxx",
	}

	ok, err := node.Add("A", "Whatever")
	if ok == false {
		t.Error("ok was false for A")
	}
	if err != nil {
		t.Error("Returned error for A")
	}

	ok, err = node.Add("B", "Hmmmm")
	if ok == false {
		t.Error("ok was false for B")
	}
	if err != nil {
		t.Error("Returned error for B")
	}

	ok, err = node.Add("C", "Hmmmm")
	if ok == false {
		t.Error("ok was false for C")
	}
	if err != nil {
		t.Error("Returned error for C")
	}

	ok, err = node.Add("D", "Hmmmm")
	if ok == false {
		t.Error("ok was false for D")
	}
	if err != nil {
		t.Error("Returned error for D")
	}

	ok, err = node.Add("E", "Hmmmm")
	if ok == false {
		t.Error("ok was false for E")
	}
	if err != nil {
		t.Error("Returned error for E")
	}

	ok, err = node.Add("G", "Hmmmm")
	if ok == false {
		t.Error("ok was false for G")
	}
	if err != nil {
		t.Error("Returned error for G")
	}

	ok, err = node.Add("H", "Hmmmm")
	if ok == false {
		t.Error("ok was false for H")
	}
	if err != nil {
		t.Error("Returned error for H")
	}

	ok, err = node.Add("I", "Hmmmm")
	if ok == false {
		t.Error("ok was false for I")
	}
	if err != nil {
		t.Error("Returned error for I")
	}

	ok, err = node.Add("J", "Hmmmm")
	if ok == false {
		t.Error("ok was false for J")
	}
	if err != nil {
		t.Error("Returned error for J")
	}

	ok, err = node.Add("A", "Whatever")
	if ok != false {
		t.Error("ok should be false for A")
	}
	if err != nil {
		t.Error("Returned error for A")
	}

	first, _ := node.GetFirstNode()
	if first.Key != "A" {
		t.Error("A is not first (" + first.Key + ")")
	}
	next := first.Next
	if next.Key != "B" {
		t.Error("B is not next (" + next.Key + ")")
	}
	next = next.Next
	if next.Key != "C" {
		t.Error("C is not next (" + next.Key + ")")
	}
	next = next.Next
	if next.Key != "D" {
		t.Error("D is not next (" + next.Key + ")")
	}
	next = next.Next
	if next.Key != "E" {
		t.Error("E is not next (" + next.Key + ")")
	}
	next = next.Next
	if next.Key != "F" {
		t.Error("F is not next (" + next.Key + ")")
	}
	next = next.Next
	if next.Key != "G" {
		t.Error("G is not next (" + next.Key + ")")
	}
	next = next.Next
	if next.Key != "H" {
		t.Error("H is not next (" + next.Key + ")")
	}
	next = next.Next
	if next.Key != "I" {
		t.Error("I is not next (" + next.Key + ")")
	}
	next = next.Next
	if next.Key != "J" {
		t.Error("J is not next (" + next.Key + ")")
	}

	match, matchErr := first.FindNode("C")
	if matchErr != nil {
		t.Error("FindNode returned an error")
	}
	if match == nil {
		t.Error("FindNode returned nil")
	} else if match.Key != "C" {
		t.Errorf("FindNode returned wrong key (%s)\n", match.Key)
	}

	match, matchErr = first.FindNode("J")
	if matchErr != nil {
		t.Error("FindNode returned an error")
	}
	if match == nil {
		t.Error("FindNode returned nil")
	} else if match.Key != "J" {
		t.Errorf("FindNode returned wrong key (%s)\n", match.Key)
	}

	match, matchErr = first.FindNode("K")
	if matchErr != nil {
		t.Error("FindNode returned an error")
	}
	if match != nil {
		t.Error("FindNode should return nil")
	}
}
