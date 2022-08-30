package primegen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollardRho(t *testing.T) {
	d := PollardRho(78_795_319 * 78_797_101)
	assert.True(t, d == 78795319 || d == 78797101)
}

func TestFactorMap(t *testing.T) {
	isPrime, p, k := FactorsMap(100)
	assert.True(t, isPrime[67])
	assert.False(t, isPrime[65])
	assert.Equal(t, []int64{2}, p[2])
	assert.Equal(t, []int64{1}, k[2])
	assert.Equal(t, []int64{2}, p[4])
	assert.Equal(t, []int64{2}, k[4])
	assert.Equal(t, []int64{3, 11}, p[99])
	assert.Equal(t, []int64{2, 1}, k[99])
}

func TestFactors(t *testing.T) {
	_, pfs := FactorsSlice(100_000)
	_, ps, as := FactorsMap(100_000)

	for i := 0; i < 10; i++ {
		fmt.Println(pfs[i], ps[int64(i)], as[int64(i)])
	}
}

func BenchmarkFactorsSlice(b *testing.B) {
	FactorsSlice(100_000_00)
}

func BenchmarkFactorsMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorsMap(100_000_00)
	}
}
