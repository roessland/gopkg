package primegen

import "testing"
import "github.com/roessland/gopkg/sliceutil"

func TestMap_low_nums(t *testing.T) {
	var m, M []bool

	m = Map(-1)
	M = []bool{}
	if !sliceutil.BoolsEqual(m, M) {
		t.Errorf("Map(-1) = %v, want %v", m, M)
	}

	m = Map(0)
	M = []bool{false}
	if !sliceutil.BoolsEqual(m, M) {
		t.Errorf("Map(0) = %v, want %v", m, M)
	}

	m = Map(1)
	M = []bool{false, false}
	if !sliceutil.BoolsEqual(m, M) {
		t.Errorf("Map(1) = %v, want %v", m, M)
	}

	m = Map(2)
	M = []bool{false, false, true}
	if !sliceutil.BoolsEqual(m, M) {
		t.Errorf("Map(2) = %v, want %v", m, M)
	}
}

func TestSlice_low_nums(t *testing.T) {
	for _, N := range []int{-1, 0, 1} {
		if p := Slice(Map(N)); !sliceutil.IntsEqual(p, []int{}) {
			t.Errorf("Slice(Map(%v)) = %v, want %v", N, p, []int{})
		}
	}
}

func TestSlice_thousand(t *testing.T) {
	theprimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
		53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131,
		137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211,
		223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293,
		307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389,
		397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479,
		487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587,
		593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673,
		677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773,
		787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881,
		883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991,
		997}
	ourprimes := Slice(Map(1000))
	if !sliceutil.IntsEqual(theprimes, ourprimes) {
		t.Errorf("Slice(Map(1000)) = %v, want %v", ourprimes, theprimes)
	}
}

func TestCount(t *testing.T) {
    isPrime := Map(1000000)
    if Count(isPrime, 10) != 4 {
        t.Errorf("Wrong number of primes below 10")
    }
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Map(10000)
	}
}

func BenchmarkSlice(b *testing.B) {
    primesMap := Map(10000)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = Slice(primesMap)
    }
}
