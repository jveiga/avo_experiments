//go:build ignore
// +build ignore

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
	// . "github.com/mmcloughlin/avo/reg"
)

func main() {
	TEXT("Add_Simple", NOSPLIT, "func(x []uint64, y []uint64) ")
	Doc("Bla adds x with y. results in x")
	x := Mem{Base: Load(Param("x").Base(), GP64())}
	y := Mem{Base: Load(Param("y").Base(), GP64())}
	n := Load(Param("x").Len(), GP64())
	s := GP64()
	XORQ(s, s)
	idx := GP64()
	XORQ(idx, idx)

	Label("loop")
	Comment("Loop until n == idx")
	CMPQ(n, idx)
	JE(LabelRef("done"))

	Comment("s =  x[i] + y[i]")
	ADDQ(x.Idx(idx, 8), s)
	ADDQ(y.Idx(idx, 8), s)
	MOVQ(s, x.Idx(idx, 8))
	XORQ(s, s)
	INCQ(idx)
	JMP(LabelRef("loop"))

	Label("done")
	// MOVQ(x, Load(Param("x"), GP64()))
	RET()
	Generate()
}
