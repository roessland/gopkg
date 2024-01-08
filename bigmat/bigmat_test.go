package bigmat_test

import (
	"math/big"
	"strings"
	"testing"

	"github.com/roessland/gopkg/bigmat"
	"github.com/stretchr/testify/require"
)

func TestBigMat_Basic(t *testing.T) {
	A := bigmat.Zeros(2, 3)
	require.Equal(t, 2, A.Rows())
	require.Equal(t, 3, A.Cols())
	require.EqualValues(t, 0, A.AtInt64(0, 0))

	// Initially 0//1
	require.EqualValues(t, 0, A.At(1, 2).Num().Int64())
	require.EqualValues(t, 1, A.At(1, 2).Denom().Int64())

	A.Set(0, 0, big.NewRat(1, 7))
	require.InDelta(t, 0.14285714285714285, A.AtFloat64(0, 0), 0.00000001)
}

func TestBigMat_String(t *testing.T) {
	A := bigmat.Zeros(2, 3)
	for k := 0; k < 6; k++ {
		(*bigmat.Vec)(A).Set(k, big.NewRat(int64(k), 1))
	}

	v := A.AsVec()

	{
		// [     0/1      1/1      2/1
		//       3/1      4/1      5/1 ]
		str := A.String()
		newLines := strings.Count(str, "\n")
		require.Equal(t, 1, newLines)
	}

	{
		// [     0/1;      1/1;      2/1;      3/1;      4/1 ]
		str := v.String()
		newLines := strings.Count(str, "\n")
		semis := strings.Count(str, ";")
		require.Equal(t, 0, newLines)
		require.Equal(t, 5, semis)
	}
}

func TestBigMat_Slice(t *testing.T) {
	A := bigmat.Zeros(2, 3)
	for k := 0; k < 6; k++ {
		(*bigmat.Vec)(A).Set(k, big.NewRat(int64(k), 1))
	}

	// a[0,:]
	row0 := A.Row(0)
	require.Equal(t, 3, row0.Len())
	require.EqualValues(t, []int{0, 1, 2}, row0.Ints())

	// a[1,:]
	row1 := A.Row(1)
	require.Equal(t, 3, row1.Len())
	require.EqualValues(t, []int{3, 4, 5}, row1.Ints())

	// a[2,:]
	require.Panics(t, func() {
		A.Row(2)
	})

	// a[:,0]
	col0 := A.Col(0)
	require.Equal(t, 2, col0.Len())
	require.EqualValues(t, []int{0, 3}, col0.Ints())

	// a[:,1]
	col1 := A.Col(1)
	require.Equal(t, 2, col1.Len())
	require.EqualValues(t, []int{1, 4}, col1.Ints())

	// a[:,2]
	col2 := A.Col(2)
	require.Equal(t, 2, col2.Len())
	require.EqualValues(t, []int{2, 5}, col2.Ints())

	// a[:,3]
	require.Panics(t, func() {
		A.Col(3)
	})
}

func TestBigMat_MatMul(t *testing.T) {
	// 0 1 2
	// 3 4 5
	A := bigmat.Zeros(2, 3)
	for k := 0; k < 6; k++ {
		(*bigmat.Vec)(A).Set(k, big.NewRat(int64(k), 1))
	}
	// 0 1
	// 2 3
	// 4 5
	B := bigmat.Zeros(3, 2)
	for k := 0; k < 6; k++ {
		(*bigmat.Vec)(B).Set(k, big.NewRat(int64(k), 1))
	}
	// 10 13
	// 28 40
	C := A.MatMul(B)
	require.Equal(t, 2, C.Rows())
	require.Equal(t, 2, C.Cols())
	require.EqualValues(t, 10, C.AtInt64(0, 0))
	require.EqualValues(t, 13, C.AtInt64(0, 1))
	require.EqualValues(t, 28, C.AtInt64(1, 0))
	require.EqualValues(t, 40, C.AtInt64(1, 1))
}

func TestBigMat_LUFact(t *testing.T) {
	// 2 2 3
	// 5 9 10
	// 4 1 2
	A := bigmat.Zeros(4, 4)
	A.SetInt64(0, 0, 2)
	A.SetInt64(0, 1, 0)
	A.SetInt64(0, 2, 4)
	A.SetInt64(0, 3, 3)
	A.SetInt64(1, 0, -4)
	A.SetInt64(1, 1, 5)
	A.SetInt64(1, 2, -7)
	A.SetInt64(1, 3, -10)
	A.SetInt64(2, 0, 1)
	A.SetInt64(2, 1, 15)
	A.SetInt64(2, 2, 2)
	A.Set(2, 3, big.NewRat(-9, 2))
	A.SetInt64(3, 0, -2)
	A.SetInt64(3, 1, 0)
	A.SetInt64(3, 2, 2)
	A.SetInt64(3, 3, 13)

	L, U := A.LUFact()

	// Check that A = L*U
	A2 := L.MatMul(U)
	res := A.Sub(A2)
	norm := res.Norm1()
	normF, _ := norm.Float64()
	require.EqualValues(t, 0, normF)
}

func TestBigMat_Norm1(t *testing.T) {
	{
		A := bigmat.Zeros(2, 3)
		norm := A.Norm1()
		normF, _ := norm.Float64()
		require.EqualValues(t, 0, normF)
	}

	{
		// 1 5 3
		// 2 6 4
		A := bigmat.Zeros(2, 3)
		A.SetInt64(0, 0, 1)
		A.SetInt64(1, 0, 2)
		A.SetInt64(0, 1, 5)
		A.SetInt64(1, 1, 6)
		A.SetInt64(0, 2, 3)
		A.SetInt64(1, 2, 4)
		norm := A.Norm1()
		require.True(t, norm.IsInt())
		normN := norm.Num().Int64()
		require.EqualValues(t, 11, normN)
	}
}

func TestBigMat_PLUFact_Nice(t *testing.T) {
	// 2 2 3
	// 5 9 10
	// 4 1 2
	A := bigmat.Zeros(4, 4)
	A.SetInt64(0, 0, 2)
	A.SetInt64(0, 1, 0)
	A.SetInt64(0, 2, 4)
	A.SetInt64(0, 3, 3)
	A.SetInt64(1, 0, -4)
	A.SetInt64(1, 1, 5)
	A.SetInt64(1, 2, -7)
	A.SetInt64(1, 3, -10)
	A.SetInt64(2, 0, 1)
	A.SetInt64(2, 1, 15)
	A.SetInt64(2, 2, 2)
	A.Set(2, 3, big.NewRat(-9, 2))
	A.SetInt64(3, 0, -2)
	A.SetInt64(3, 1, 0)
	A.SetInt64(3, 2, 2)
	A.SetInt64(3, 3, 13)

	L, U := A.LUFact()

	// Check that A = L*U
	A2 := L.MatMul(U)
	res := A.Sub(A2)
	norm := res.Norm1()
	normF, _ := norm.Float64()
	require.EqualValues(t, 0, normF)
}

func TestBigMat_VecAbs(t *testing.T) {
	u := bigmat.ZerosVec(4)
	u.SetInt64(0, 2)
	u.SetInt64(1, -2)
	u.SetInt64(2, 1)
	u.SetInt64(3, -4)

	absU := u.Abs()

	require.EqualValues(t, 2, u.AtInt64(0))
	require.EqualValues(t, -2, u.AtInt64(1))
	require.EqualValues(t, 1, u.AtInt64(2))
	require.EqualValues(t, -4, u.AtInt64(3))

	require.EqualValues(t, 2, absU.AtInt64(0))
	require.EqualValues(t, 2, absU.AtInt64(1))
	require.EqualValues(t, 1, absU.AtInt64(2))
	require.EqualValues(t, 4, absU.AtInt64(3))
}

func TestBigMag_Argmax(t *testing.T) {
	u := bigmat.ZerosVec(4)
	u.SetInt64(0, 2)
	u.SetInt64(1, -2)
	u.SetInt64(2, 1)
	u.SetInt64(3, -4)

	argmax := u.Argmax(func(r *big.Rat) *big.Rat { return new(big.Rat).Abs(r) })
	require.Equal(t, 3, argmax)
}

func TestBigMat_PLUFact_Requires_Pivoting(t *testing.T) {
	// 2 0 4 3
	// -2 0 2 13
	// 1 15 2 -4
	// -4 5 -7 10
	A := bigmat.Zeros(4, 4)
	A.SetInt64(0, 0, 2)
	A.SetInt64(0, 1, 0)
	A.SetInt64(0, 2, 4)
	A.SetInt64(0, 3, 3)
	A.SetInt64(1, 0, -2)
	A.SetInt64(1, 1, 0)
	A.SetInt64(1, 2, 2)
	A.SetInt64(1, 3, -13)
	A.SetInt64(2, 0, 1)
	A.SetInt64(2, 1, 15)
	A.SetInt64(2, 2, 2)
	A.Set(2, 3, big.NewRat(-9, 2))
	A.SetInt64(3, 0, -4)
	A.SetInt64(3, 1, 5)
	A.SetInt64(3, 2, -7)
	A.SetInt64(3, 3, -10)

	selectPivot := func(k int) int {
		pivot := A.Col(k).Argmax(func(r *big.Rat) *big.Rat { return new(big.Rat).Abs(r) })
		return pivot
	}

	require.Equal(t, 3, selectPivot(0))

	L, U, p := A.PLUFact()

	// Check that ~A = L*U
	tildeA := A.Pivot(p) // Permute rows
	LU := L.MatMul(U)
	res := tildeA.Sub(LU)
	norm := res.Norm1()
	normF, _ := norm.Float64()
	require.EqualValues(t, 0, normF)
}

func TestBigMat_NewMatFromString(t *testing.T) {
	oneHalf, ok := new(big.Rat).SetString("1/2")
	require.True(t, ok)
	require.EqualValues(t, 1, oneHalf.Num().Int64())

	A, err := bigmat.NewMatFromString(`1337//1338`)
	require.NoError(t, err)
	require.EqualValues(t, 1337, A.At(0, 0).Num().Int64())
	require.EqualValues(t, 1338, A.At(0, 0).Denom().Int64())

	u := bigmat.MustNewVecFromString(`
    1337//1338
    1/1000
  `)
	require.EqualValues(t, 2, u.Len())
	require.EqualValues(t, 1337, u.At(0).Num().Int64())
	require.EqualValues(t, 1000, u.At(1).Denom().Int64())
}

func TestBigMat_PLUFact_AoC2023_24(t *testing.T) {
	A, err := bigmat.NewMatFromString(`
    0    517   430  0                200449925047571 119796160238546
    -517 0    -120 -200449925047571  0               -50016083636040
    -430 120  0    -119796160238546  50016083636040  0
    0    264  217   0                5508734318078  -43844419026563
    -264 0   -189  -5508734318078    0              -103223937923385
    -217 189 0     43844419026563    103223937923385 0
  `)
	require.NoError(t, err)

	L, U, p := A.PLUFact()

	LU := L.MatMul(U)
	tildeA := A.Pivot(p) // Permute rows
	res := tildeA.Sub(LU)
	norm := res.Norm1()
	require.True(t, norm.Cmp(big.NewRat(0, 1)) == 0)
}

func TestBigMat_Forwardsub(t *testing.T) {
	L := bigmat.MustNewMatFromString(`
    1//1         0//1              0//1                     0//1            0//1
    33//488       1//1              0//1                     0//1            0//1
  597//976  -55749//792862         1//1                     0//1            0//1
  245//976  413467//792862  10083539//80440138              1//1            0//1
  177//488   34161//56633   43183462//120660207  -34128982256//61666927185  1//1
  `)

	b := bigmat.MustNewVecFromString(`
      670//1
      131//1
      925//1
      826//1
      823//1
    `)

	xActual := L.Forwardsub(b)
	xExpected := bigmat.MustNewVecFromString(`
             670//1
          20909//244
       59034187//113266
    88128801929//160880276
 13257035041413//20555642395
  `)

	require.True(t, xExpected.Equal(xActual))
}

func TestBigMat_Backsub(t *testing.T) {
	U := bigmat.MustNewMatFromString(`
 976//1     113//1          350//1                895//1                   478//1
   0//1  396431//488      75721//244           412593//488              226109//244
   0//1       0//1    241320414//396431    -131220766//396431        535003271//792862
   0//1       0//1            0//1       -20555642395//40220069   -43892193245//160880276
   0//1       0//1            0//1                  0//1         -101762259857//141763051
  `)

	b := bigmat.MustNewVecFromString(`
      670//1
      131//1
      925//1
      826//1
      823//1
    `)

	xActual := U.Backsub(b)
	xExpected := bigmat.MustNewVecFromString(`
 19077382127791630934866447173755692177//16276013367835250878062971351401069480
   221033844124556618912751507666463623//133409945637993859656253863536074340
          13022571089199006066794148891//5802198810176639847324557830
                -8401946731626362742047//8367154491710223350060
                          -116670990973//101762259857
  `)

	require.True(t, xExpected.Equal(xActual))
}

func TestBigMat_Backslash(t *testing.T) {
	A := bigmat.MustNewMatFromString(`
     66//1  820//1  334//1  906//1  959//1
    245//1  452//1  326//1  113//1  415//1
    976//1  113//1  350//1  895//1  478//1
    597//1   12//1  801//1  157//1  902//1
    354//1  531//1  532//1  999//1  407//1
  `)

	b := bigmat.MustNewVecFromString(`
      670//1
      131//1
      925//1
      826//1
      823//1
    `)

	xActual := A.Backslash(b)
	xExpected := bigmat.MustNewVecFromString(`
        974070714041//88533166075590
     -16912724457027//29511055358530
      22828392056381//44266583037795
      12640691387489//17706633215118
      14779216313252//44266583037795
  `)

	require.True(t, xExpected.Equal(xActual))
}
