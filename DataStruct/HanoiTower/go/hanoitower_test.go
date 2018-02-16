// Copyright 2018 Itamar Ostricher

package hanoitower

import "testing"

func TestHanoiTower(t *testing.T) {
	var hn HanoiTower
	if s := hn.Size(); s != 0 {
		t.Errorf("Expected new HN to have size 0, got %d", s)
	}
	hn.Push(10)
	if s := hn.Size(); s != 1 {
		t.Errorf("Expected new HN to have size 1, got %d", s)
	}
	hn.Push(5)
	hn.Push(1)
	if s := hn.Size(); s != 3 {
		t.Errorf("Expected new HN to have size 3, got %d", s)
	}
	if hn.top.value != 1 || hn.top.next.value != 5 || hn.top.next.next.value != 10 {
		t.Errorf("Expected Tower of 10-5-1, got something else")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected push(7) to panic, but it didn't")
		}
	}()
	hn.Push(7)
}
