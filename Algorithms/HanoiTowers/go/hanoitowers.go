// Copyright 2018 Itamar Ostricher
// Recursive algorithm for solving the Hanoi Towers riddle

package main

import (
	"fmt"
	hanoitower "github.com/itamaro/clrs/DataStruct/HanoiTower/go"
)

func main() {
	var a, b, c hanoitower.HanoiTower
	for i := 11; i > 0; i -= 2 {
		a.Push(i)
	}
	fmt.Println(" Starting position")
	PrintAll(&a, &b, &c)
	SolveHanoi(&a, &b, &c)
	fmt.Println("  Ending position")
	PrintAll(&a, &b, &c)
}

func SolveHanoi(src, dst, swap *hanoitower.HanoiTower) {
	if src.Size() == 0 {
		return
	}
	moveRings(src, dst, swap, src.Size())
}

func moveRings(src, dst, swap *hanoitower.HanoiTower, n int) {
	if n == 1 {
		dst.Push(src.Pop())
		return
	}
	moveRings(src, swap, dst, n - 1)
	dst.Push(src.Pop())
	moveRings(swap, dst, src, n - 1)
}

func PrintAll(a, b, c *hanoitower.HanoiTower) {
	fmt.Println("====================")
	fmt.Println("^^^^^^^^^ a ^^^^^^^^")
	a.Print()
	fmt.Println("^^^^^^^^^ b ^^^^^^^^")
	b.Print()
	fmt.Println("^^^^^^^^^ c ^^^^^^^^")
	c.Print()
	fmt.Println("====================")
	fmt.Println("")
}
