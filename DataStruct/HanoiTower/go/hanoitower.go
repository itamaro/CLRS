// Copyright 2018 Itamar Ostricher
// hanoitower package providing a "custom stack" describing a Hanoi tower
// (a stack that can only hold "rings" with decreasing "radius")
// for educational purposes

package hanoitower

import (
	"fmt"
	"strings"
)

// HanoiTower stack struct provides stack semantics data structure,
// with O(1) push/pop/size operations
// and enforces the Hanoi tower semantics on insertion
type HanoiTower struct {
	top *ring
	size int
}

// internal structure for stack nodes ("rings")
type ring struct {
	next *ring
	value int
}

// Push a new ring to the top of the tower
// Runs in O(1) with no extra memory
// Panics if the Hanoi invariant is violated
func (hn *HanoiTower) Push(value int) {
	if hn.top != nil {
		if value >= hn.top.value {
			panic("Hanoi violation")
		}
	}
	newRing := &ring{hn.top, value}
	hn.top = newRing
	hn.size++
}

// Pop the top item from the stack
// Runs in O(1) with no extra memory
func (hn *HanoiTower) Pop() int {
	if hn.top != nil {
		value := hn.top.value
		hn.top = hn.top.next
		hn.size--
		return value
	}
	panic("Popping empty HanoiTower")
}

// Return the number of rings in the tower
// Runs in O(1) with no extra memory
func (hn HanoiTower) Size() int {
	return hn.size
}

// Print the Hanoi tower using '*' for rings (works with up to diameter=20)
func (hn HanoiTower) Print() {
	for r := hn.top; r != nil; r = r.next {
		fmt.Printf("%s%s\n", strings.Repeat(" ", 10 - r.value / 2), strings.Repeat("*", r.value))
	}
}

func printHelper(r ring) int {
	if r.next == nil {
			return r.value
	}
	width := printHelper(*r.next)
	fmt.Printf("%s\n", "*")
	return width
}
