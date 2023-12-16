package crt_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/roessland/gopkg/mathutil/crt"
	"github.com/stretchr/testify/assert"
)

func TestCRT(t *testing.T) {
	t.Run("is correct for 1 equation", func(t *testing.T) {
		{
			// x = 1 (mod 2)
			// x = 1
			// N = 2
			x, N := crt.CRT([]int64{1}, []int64{2})

			assert.EqualValues(t, x, 1)
			assert.EqualValues(t, N, 2)
		}
	})

	t.Run("is correct for 2 equations", func(t *testing.T) {
		{
			// x = 2 (mod 3)
			// x = 3 (mod 4)
			x, N := crt.CRT([]int64{2, 3}, []int64{3, 4})
			assert.EqualValues(t, 11, x)
			assert.EqualValues(t, 12, N)
			assert.EqualValues(t, 2, x%3)
			assert.EqualValues(t, 3, x%4)
		}
	})

	t.Run("is correct for 3 equations", func(t *testing.T) {
		{
			// x = 3 (mod 37)
			// x = 7 (mod 29)
			// x = 1 (mod 3)
			// x ≡ 2704 (mod 3219).
			x, N := crt.CRT([]int64{3, 7, 1}, []int64{37, 29, 3})
			assert.EqualValues(t, 2704, x)
			assert.EqualValues(t, 3219, N)
			assert.EqualValues(t, 3, x%37)
			assert.EqualValues(t, 7, x%29)
			assert.EqualValues(t, 1, x%3)
		}
	})

	t.Run("is correct for 2 equations, #2", func(t *testing.T) {
		{
			// x = 32 (mod 35)
			// x = 2 (mod 27)
			x, N := crt.CRT([]int64{32, 2}, []int64{35, 27})
			assert.EqualValues(t, 137, x)
			assert.EqualValues(t, 945, N)
			assert.EqualValues(t, 32, x%35)
			assert.EqualValues(t, 2, x%27)
		}
	})

	t.Run("panics if ns aren't pairwise coprime", func(t *testing.T) {
		{
			// This system of equations has a valid solution, but it's not in the
			// format required by CRT.
			// x = 3 (mod 74)
			// x = 3 (mod 37)
			// x = 7 (mod 29)
			// x = 1 (mod 3)
			// x ≡ 2704 (mod 3219).
			assert.Panics(t, func() {
				crt.CRT([]int64{3, 3, 7, 1}, []int64{74, 37, 29, 3})
			})
		}
	})

	t.Run("panics if any ns are <= 1", func(t *testing.T) {
		{
			assert.Panics(t, func() {
				crt.CRT([]int64{3, 3, 0}, []int64{37, 29, 1})
			})
		}
	})
}

// If CRT doesn't panic, the solution should be valid.
func FuzzCRT_2_equations_basic(f *testing.F) {
	tryCRT := func(as []int64, ns []int64) (x int64, N int64, err error) {
		defer func() {
			if r := recover(); r != nil {
				x, N = 0, 0
				str := fmt.Sprint(r)
				if !strings.Contains(str, "must be") && !strings.Contains(str, "to ensure") && !strings.Contains(str, "pairwise coprime") {
					panic(r)
				}
				err = errors.New(fmt.Sprint(r))
			}
		}()
		x, N = crt.CRT(as, ns)
		return x, N, nil
	}

	f.Add(int64(1), int64(2), int64(3), int64(4))

	f.Fuzz(func(t *testing.T, a1, a2, n1, n2 int64) {
		as := []int64{a1, a2}
		ns := []int64{n1, n2}
		x, N, err := tryCRT(as, ns)
		if err != nil {
			t.Skip()
		}
		if x < 0 {
			t.Errorf("x = %d (must be positive)", x)
		}
		if N <= 0 {
			t.Errorf("N = %d (must be posiive)", N)
		}
		if x >= N {
			t.Errorf("x = %d, N = %d (must be x < N)", x, N)
		}
		if x%n1 != a1%n1 || x%n2 != a2%n2 {
			t.Errorf("\nFor equations \n[x=%d (mod %d), x=%d (mod %d)],\ngot solution %d mod %d, but \n\t x %% %d = %d, and \n\t x %% %d = %d ",
				a1, n1, a2, n2, x, N, n1, x%n1, n2, x%n2)
			panic("yeahhhhh")
		}
	})
}

func TestCRTCoerce(t *testing.T) {
	t.Run("no valid sols #1", func(t *testing.T) {
		assert.Panics(t, func() {
			crt.CRTCoerce([]int64{3, 7}, []int64{5, 5})
		})
	})

	t.Run("is correct for 4 equations (1 coprime, bigger)", func(t *testing.T) {
		{
			// x = 3 (mod 74)
			// x = 3 (mod 37)
			// x = 7 (mod 29)
			// x = 1 (mod 3)
			//
			// Has solution:
			// x = 5923 (mod 6438)
			//
			// The system isn't valid for the normal CRT, as some factors aren't
			// coprime, but we can coerce it into a valid system:
			//
			// x = 1 (mod 2)
			// x = 3 (mod 37)
			// x = 7 (mod 29)
			// x = 1 (mod 3)
			//
			// With solution:
			// x ≡ 5923 (mod 6438).
			x_, N_ := crt.CRT([]int64{1, 3, 7, 1}, []int64{2, 37, 29, 3})
			assert.EqualValues(t, 5923, x_)
			assert.EqualValues(t, 6438, N_)

			assert.EqualValues(t, 3, x_%74)
			assert.EqualValues(t, 3, x_%37)
			assert.EqualValues(t, 7, x_%29)
			assert.EqualValues(t, 1, x_%3)

			assert.EqualValues(t, 3, (x_+N_)%74)
			assert.EqualValues(t, 3, (x_+N_)%37)
			assert.EqualValues(t, 7, (x_+N_)%29)
			assert.EqualValues(t, 1, (x_+N_)%3)

			x, N := crt.CRTCoerce([]int64{3, 3, 7, 1}, []int64{74, 37, 29, 3})
			assert.EqualValues(t, 5923, x)
			assert.EqualValues(t, 6438, N)
			assert.EqualValues(t, 3, x%74)
			assert.EqualValues(t, 3, x%37)
			assert.EqualValues(t, 7, x%29)
			assert.EqualValues(t, 1, x%3)
		}
	})

	t.Run("panics for system without solutions", func(t *testing.T) {
		// Has no solutions:
		// x = 2 (mod 74)
		// x = 3 (mod 37)
		//
		// This has a solution though:
		// x = 2 (mod 2)
		// x = 3 (mod 37)
		assert.Panics(t, func() {
			fmt.Println(crt.CRTCoerce([]int64{2, 3}, []int64{74, 37}))
		}, "should have failed, cannot have x=2(mod 74) and x=3(mod 37) at the same time")
	})

	t.Run("allows non-simplified input", func(t *testing.T) {
		t.Skip()
		// x = 77 (mod 74)
		// x = 40 (mod 37)
		x, N := crt.CRTCoerce([]int64{77, 40}, []int64{74, 37})
		assert.EqualValues(t, 3, x%74)
		assert.EqualValues(t, 3, x%37)
		assert.EqualValues(t, 3, (x+N)%74)
		assert.EqualValues(t, 3, (x+N)%37)
	})

	t.Run("allows over-determined system (3 equations, 1 redundant)", func(t *testing.T) {
		t.Skip()
		// x = 3 (mod 37)
		// x = 3 (mod 37)
		// x = 5 (mod 7)
		x, N := crt.CRTCoerce([]int64{3, 3, 5}, []int64{37, 37, 7})
		assert.EqualValues(t, 3, x%74)
		assert.EqualValues(t, 3, x%37)
		assert.EqualValues(t, 3, (x+N)%74)
		assert.EqualValues(t, 3, (x+N)%37)
	})

	t.Run("negative a", func(t *testing.T) {
		// x = -108 (mod 175)
		// x = 2 (mod 135)
		x, N := crt.CRTCoerce([]int64{-108, 2}, []int64{175, 135})
		t.Log(x, N)
		assert.EqualValues(t, -108+175, x%175)
		assert.EqualValues(t, 2, x%135)
		assert.EqualValues(t, 4725, N)
	})

	t.Run("example 2", func(t *testing.T) {
		x, N := crt.CRTCoerce([]int64{23, 11}, []int64{27, 102})
		t.Log(x, N)
		assert.EqualValues(t, 725, x)
		assert.EqualValues(t, 918, N)
	})

	t.Run("example 3", func(t *testing.T) {
		// 98 = 2 * 7 * 7
		// 224 = 2 * 2 * 2 * 2 * 2 * 7
		x, N := crt.CRTCoerce([]int64{88, 130}, []int64{98, 224})
		t.Log(x, N)
		assert.EqualValues(t, 578, x)
		assert.EqualValues(t, 1568, N)
	})

	t.Run("example 4", func(t *testing.T) {
		// 16 = 2 * 2 * 2 * 2
		// 6 = 2 * 3
		// x=1 (mod 16)
		// x=2 (mod 6)
		//
		// x=3 (mod 48)
		// x=8 (mod 48) (conflict!)
		//
		// Figure this out by
		// Multiply equation i by lcm/n_i
		assert.Panics(t, func() {
			crt.CRTCoerce([]int64{1, 2}, []int64{16, 6})
		})
	})
}

func tryCRTCoerce(as []int64, ns []int64) (x int64, N int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			x, N = 0, 0
			str := fmt.Sprint(r)
			if !strings.Contains(str, "simultaneously") &&
				!strings.Contains(str, "must be pos") &&
				!strings.Contains(str, "cannot be neg") &&
				!strings.Contains(str, "no solution") &&
				!strings.Contains(str, "must be greater") {
				panic(r)
			}
			err = errors.New(fmt.Sprint(r))
		}
	}()
	x, N = crt.CRTCoerce(as, ns)
	return x, N, nil
}

// If CRTCoerce doesn't panic, the solution should be valid.
func FuzzCRTCoerce2_equations_basic(f *testing.F) {
	f.Add(int64(1), int64(2), int64(3), int64(4))

	f.Fuzz(func(t *testing.T, a1, a2, n1, n2 int64) {
		as := []int64{a1, a2}
		ns := []int64{n1, n2}
		t.Log(as, ns)
		x, N, err := tryCRTCoerce(as, ns)
		if err != nil {
			t.Skip()
		}
		if x < 0 {
			t.Errorf("x = %d (must be positive)", x)
		}
		if N <= 0 {
			t.Errorf("N = %d (must be posiive)", N)
		}
		if x >= N {
			t.Errorf("x = %d, N = %d (must be x < N)", x, N)
		}
		t.Log("Got x, N", x, N)
		for a1 < 0 {
			a1 += n1
		}
		for a2 < 0 {
			a2 += n2
		}
		t.Log("Eq 1: x =", a1, "mod", n1)
		t.Log("Eq 2: x =", a2, "mod", n2)
		if x%n1 != a1%n1 || x%n2 != a2%n2 {

			t.Log("x%n1", x%n1, "a1%n1", a1%n1, "x%n2", x%n2, "a2%n2", a2%n2)
			t.Errorf("\nFor equations \n[x=%d (mod %d), x=%d (mod %d)],\ngot solution %d mod %d, but \n\t x %% %d = %d, and \n\t x %% %d = %d ",
				a1, n1, a2, n2, x, N, n1, x%n1, n2, x%n2)
		}
	})
}

// If CRTCoerce doesn't panic, the solution should be valid.
func FuzzCRTCoerce3_equations_basic(f *testing.F) {
	f.Add(int64(1), int64(2), int64(3), int64(4), int64(7), int64(9))

	f.Fuzz(func(t *testing.T, a1, a2, a3, n1, n2, n3 int64) {
		as := []int64{a1, a2, a3}
		ns := []int64{n1, n2, n3}
		t.Log(as, ns)
		x, N, err := tryCRTCoerce(as, ns)
		if err != nil {
			t.Skip()
		}
		if x < 0 {
			t.Errorf("x = %d (must be positive)", x)
		}
		if N <= 0 {
			t.Errorf("N = %d (must be posiive)", N)
		}
		if x >= N {
			t.Errorf("x = %d, N = %d (must be x < N)", x, N)
		}
		t.Log("Got x, N", x, N)
		for a1 < 0 {
			a1 += n1
		}
		for a2 < 0 {
			a2 += n2
		}
		for a3 < 0 {
			a3 += n3
		}
		t.Log("Eq 1: x =", a1, "mod", n1)
		t.Log("Eq 2: x =", a2, "mod", n2)
		t.Log("Eq 3: x =", a3, "mod", n3)
		if x%n1 != a1%n1 || x%n2 != a2%n2 || x%n3 != a3%n3 {

			t.Log("x%n1", x%n1, "a1%n1", a1%n1, "x%n2", x%n2, "a2%n2", a2%n2, "x%n3", x%n3, "a3%n3", a3%n3)
			t.Errorf("\nFor equations \n[x=%d (mod %d), x=%d (mod %d), x=%d (mod %d)],\ngot solution %d mod %d, but \n\t x %% %d = %d, and \n\t x %% %d = %d, and \n\t x %% %d = %d",
				a1, n1, a2, n2, a3, n3, x, N, n1, x%n1, n2, x%n2, n3, x%n3)
		}
	})
}
