package tests_test

import (
	EComBase "github.com/codedv8/go-ecom-base"
	"math/rand"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestListSize(t *testing.T) {
	n := 17
	size := (n + 1) / 2
	if size != 9 {
		t.Errorf("Size for %d should be 9 but we got %d\n", n, size)
	}

	n = 457
	size = (n + 1) / 2
	if size != 229 {
		t.Errorf("Size for %d should be 229 but we got %d\n", n, size)
	}

	s := "12345*"
	t.Logf("Last char would be %s", s[len(s)-1:])
}

func TestLargeSetBalance(t *testing.T) {
	tree := &EComBase.LinkedTree{}
	n := 0
	start := time.Now()
	for i := 0; i < 129418; i++ {
		rand.Seed(time.Now().UnixNano())
		chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
			"abcdefghijklmnopqrstuvwxyzåäö" +
			"0123456789")
		length := 48
		var b strings.Builder
		for i := 0; i < length; i++ {
			b.WriteRune(chars[rand.Intn(len(chars))])
		}
		str := b.String() // E.g. "ExcbsVQs"
		ok, _ := tree.Add(str, nil)
		if ok == true {
			n++
		}
	}
	t.Logf("Adding %d nodes took %f s\n", n, time.Since(start).Seconds())
	start = time.Now()
	tree.Balance()
	t.Logf("Balance took %f s\n", time.Since(start).Seconds())
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	t.Logf("Memory Alloc %d KiB\n", m.Alloc/1024)
	t.Logf("Sys %d KiB\n", m.Sys/1024)
}

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
	tree.Add("Charlie", nil)
	tree.Add("Char*", nil)

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

	find, findErr := tree.FindNode("Charlie")
	if findErr != nil {
		t.Error("Find failed")
	} else if find == nil {
		t.Error("Find should not return nil if successful")
	} else {
		t.Logf("We found %s\n", find.Key)
	}
	find, findErr = tree.FindNode("Charlie Watts")
	if findErr != nil {
		t.Error("Find failed")
	} else if find == nil {
		t.Error("Find should not return nil if successful")
	} else {
		t.Logf("We found %s\n", find.Key)
	}
	find, findErr = tree.FindNode("Charli")
	if findErr != nil {
		t.Error("Find failed")
	} else if find == nil {
		t.Error("Find should not return nil if successful")
	} else {
		t.Logf("We found %s\n", find.Key)
	}
}
