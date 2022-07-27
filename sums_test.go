package avo_experiments

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddSimple(t *testing.T) {
	size := uint(10)
	x, y := random_pair_uint_array(size)
	Add_Simple(x, y)
	nx := make([]uint64, size)
	copy(x, nx)
	require.Equal(t, x, nx)
}

func BenchmarkAddSimple(b *testing.B) {
	size := uint(10)
	x, y := random_pair_uint_array(size)
	Add_Simple(x, y)
}

func BenchmarkAdd(b *testing.B) {
	size := uint(10)
	x, y := random_pair_uint_array(size)
	Add(x, y)
}

func FuzzAddSimple(f *testing.F) {
	rand.Seed(time.Now().Unix())
	f.Fuzz(func(t *testing.T, size uint) {
		x, y := random_pair_uint_array(size)
		nx := make([]uint64, size)
		copy(x, nx)
		Add_Simple(x, y)
		Add(nx, y)
		require.Equal(t, x, nx)
	})
}

// function adds two arrays of uint64s and stores result in first array.
func Add(x []uint64, y []uint64) {
	for i := range x {
		x[i] += y[i]
	}
}

func random_pair_uint_array(size uint) ([]uint64, []uint64) {
	x := make([]uint64, size)
	y := make([]uint64, size)
	for i := range x {
		x[i] = rand.Uint64()
	}
	for i := range y {
		y[i] = rand.Uint64()
	}
	return x, y
}
