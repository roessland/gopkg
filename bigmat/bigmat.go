package bigmat

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
)

var (
	zero = big.NewRat(0, 1)
	one  = big.NewRat(1, 1)
)

// Mat is a matrix of arbitrary precision rational values.
type Mat struct {
	data []*big.Rat
	rows int
	cols int
}

// Vec is a vector of arbitrary precision rational values.
// Equivalent to a 1-column matrix.
type Vec Mat

// Zeros returns a new matrix of the given size, filled with zeros.
func Zeros(rows, cols int) *Mat {
	data_ := make([]big.Rat, rows*cols)
	data := make([]*big.Rat, rows*cols)
	for i := 0; i < rows*cols; i++ {
		data[i] = &data_[i]
	}
	return &Mat{data: data, rows: rows, cols: cols}
}

// ZerosVec returns a new vector of the given size, filled with zeros.
func ZerosVec(n int) *Vec {
	return (*Vec)(Zeros(n, 1))
}

// FillVec returns a new vector of the given size, filled with the given value.
func FillVec(n int, r *big.Rat) *Vec {
	u := ZerosVec(n)
	for i := 0; i < n; i++ {
		u.Set(i, r)
	}
	return u
}

// NewMatFromString returns a new matrix from a string representation.
func NewMatFromString(s string) (*Mat, error) {
	s = strings.Trim(s, "\n\r\t ")
	numStrs := [][]string{}
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, "//", "/")
		if line == "" {
			continue
		}
		lineStrs := strings.Fields(line)
		numStrs = append(numStrs, lineStrs)
	}
	rows := len(numStrs)
	cols := len(numStrs[0])

	A := Zeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			aij, ok := new(big.Rat).SetString(numStrs[i][j])
			if !ok {
				return nil, fmt.Errorf("could not parse %v", numStrs[i][j])
			}
			A.Set(i, j, aij)
		}
	}
	return A, nil
}

func NewVecFromString(s string) (*Vec, error) {
	m, err := NewMatFromString(s)
	if err != nil {
		return nil, err
	}
	return m.AsVec(), nil
}

func MustNewMatFromString(s string) *Mat {
	A, err := NewMatFromString(s)
	if err != nil {
		panic(err)
	}
	return A
}

func MustNewVecFromString(s string) *Vec {
	v, err := NewVecFromString(s)
	if err != nil {
		panic(err)
	}
	if v == nil {
		panic("v is nil")
	}
	return v
}

// Rows returns the number of rows in the matrix.
func (A *Mat) Rows() int {
	return A.rows
}

// Len returns the number of rows in the vector.
func (v *Vec) Len() int {
	return v.rows
}

// Cols returns the number of columns in the matrix.
func (A *Mat) Cols() int {
	return A.cols
}

// At returns the value of the matrix at row i and column j.
func (A *Mat) At(i, j int) *big.Rat {
	return A.data[i*A.cols+j]
}

// At returns the value of the vector at row i.
func (v *Vec) At(i int) *big.Rat {
	return v.data[i]
}

// AsVec returns a view of the matrix as a vector.
func (A *Mat) AsVec() *Vec {
	if A == nil {
		panic("A is nil")
	}
	return &Vec{data: A.data, rows: A.rows * A.cols, cols: 1}
}

// AsCol returns a view of the matrix as a column vector.
func (v *Vec) AsCol() *Mat {
	return &Mat{data: v.data, rows: v.rows, cols: 1}
}

// AsRow returns a view of the matrix as a row vector.
func (v *Vec) AsRow() *Mat {
	return &Mat{data: v.data, rows: 1, cols: v.rows}
}

// String returns a string representation of the matrix.
func (A *Mat) String() string {
	var s bytes.Buffer
	fmt.Fprintf(&s, "[")
	for i := 0; i < A.rows; i++ {
		if i > 0 {
			fmt.Fprintf(&s, " ")
		}
		for j := 0; j < A.cols; j++ {
			fmt.Fprintf(&s, "%8v ", A.At(i, j))
		}
		if i < A.rows-1 {
			fmt.Fprintf(&s, "\n")
		}
	}
	fmt.Fprintf(&s, "]")
	return s.String()
}

// String returns a string representation of the vector.
func (v *Vec) String() string {
	var s bytes.Buffer
	fmt.Fprintf(&s, "[")
	for i := 0; i < v.rows; i++ {
		if i > 0 {
			fmt.Fprintf(&s, " ")
		}
		fmt.Fprintf(&s, "%8v", v.At(i))
		if i < v.rows-1 {
			fmt.Fprintf(&s, ";")
		}
	}
	fmt.Fprintf(&s, " ]")
	return s.String()
}

// AtInt64 returns the value of the matrix at row i and column j as an int64.
func (A *Mat) AtInt64(i, j int) int64 {
	r := A.At(i, j)
	if !r.IsInt() {
		msg := fmt.Sprintf("not an integer: %v", r)
		panic(msg)
	}
	return A.At(i, j).Num().Int64()
}

// AtInt64 returns the value of the vector at row i as an int64.
func (v *Vec) AtInt64(i int) int64 {
	return (*Mat)(v).AtInt64(i, 0)
}

// AtFloat64 returns the value of the matrix at row i and column j as a float64.
func (A *Mat) AtFloat64(i, j int) float64 {
	r, _ := A.At(i, j).Float64() // _ means "is exact?", not an error
	return r
}

// AtFloat64 returns the value of the vector at row i as a float64.
func (v *Vec) AtFloat64(i int) float64 {
	return (*Mat)(v).AtFloat64(i, 0)
}

// Int64s returns a copy of the matrix as a slice of int64s.
// If any value is not an integer, panics.
func (v *Vec) Int64s() []int64 {
	ns := make([]int64, v.rows)
	for i := 0; i < v.rows; i++ {
		ns[i] = v.AtInt64(i)
	}
	return ns
}

func (v *Vec) Ints() []int {
	ns := make([]int, v.rows)
	for i := 0; i < v.rows; i++ {
		ns[i] = int(v.AtInt64(i))
	}
	return ns
}

// Float64s returns a copy of the vector as a slice of float64s.
func (v *Vec) Float64s() []float64 {
	fs := make([]float64, v.rows)
	for i := 0; i < v.rows; i++ {
		fs[i] = v.AtFloat64(i)
	}
	return fs
}

// Row returns a view of ith row of the matrix, as a vector.
func (A *Mat) Row(i int) *Vec {
	return &Vec{data: A.data[i*A.cols : (i+1)*A.cols], rows: A.cols, cols: 1}
}

// Col returns a view of ith column of the matrix, as a vector.
func (A *Mat) Col(j int) *Vec {
	v := &Vec{data: make([]*big.Rat, A.rows), rows: A.rows, cols: 1}
	for i := 0; i < A.rows; i++ {
		v.data[i] = A.At(i, j)
	}
	return v
}

// Set is equivalent to At(i, j).Set(v)
func (A *Mat) Set(i, j int, v *big.Rat) {
	A.data[i*A.cols+j].Set(v)
}

// Set is equivalent to At(i).Set(v)
func (v *Vec) Set(i int, r *big.Rat) {
	v.data[i].Set(r)
}

// SetInt64 is equivalent to At(i, j).SetInt64(n)
func (A *Mat) SetInt64(i, j int, n int64) *big.Rat {
	A.data[i*A.cols+j].SetInt64(n)
	return A.data[i*A.cols+j]
}

// SetInt64 is equivalent to At(i).SetInt64(n)
func (v *Vec) SetInt64(i int, n int64) {
	v.data[i].SetInt64(n)
}

// SetInt sets the ith element of the vector to n.
func (v *Vec) SetInt(i int, n int) {
	v.data[i].SetInt64(int64(n))
}

// SetFloat64 is equivalent to At(i, j).SetFloat64(v)
func (A *Mat) SetFloat64(i, j int, v float64) {
	A.data[i*A.cols+j].SetFloat64(v)
}

// SetFloat64 is equivalent to At(i).SetFloat64(v)
func (v *Vec) SetFloat64(i int, n float64) {
	v.data[i].SetFloat64(n)
}

func (A *Mat) SetRow(i int, v *Vec) {
	for j := 0; j < A.cols; j++ {
		A.Set(i, j, v.At(j))
	}
}

func (A *Mat) SetCol(j int, v *Vec) {
	for i := 0; i < A.rows; i++ {
		A.Set(i, j, v.At(i))
	}
}

func (A *Mat) Copy() *Mat {
	B := Zeros(A.rows, A.cols)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			B.Set(i, j, A.At(i, j))
		}
	}
	return B
}

// Outer product of two vectors.
//
// See https://en.wikipedia.org/wiki/Outer_product .
//
// $$ (u \otimes v)_{ij} = u_i v_j $$
func (u *Vec) Outer(v *Vec) *Mat {
	A := Zeros(u.rows, v.rows)
	for i := 0; i < u.rows; i++ {
		for j := 0; j < v.rows; j++ {
			A.Set(i, j, new(big.Rat).Mul(u.At(i), v.At(j)))
		}
	}
	return A
}

func (A *Mat) Sub(B *Mat) *Mat {
	C := Zeros(A.rows, A.cols)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			C.Set(i, j, new(big.Rat).Sub(A.At(i, j), B.At(i, j)))
		}
	}
	return C
}

// Norm1 returns the 1-norm of the matrix.
// It is equivalent to the maximum absolute column sum.
func (A *Mat) Norm1() *big.Rat {
	norm := big.NewRat(0, 1)
	for j := 0; j < A.cols; j++ {
		colSum := big.NewRat(0, 1)
		for i := 0; i < A.rows; i++ {
			colSum.Add(colSum, new(big.Rat).Abs(A.At(i, j)))
		}
		if colSum.Cmp(norm) > 0 {
			norm.Set(colSum)
		}
	}
	return norm
}

// IsZero returns true if all elements of the matrix are zero.
func (A *Mat) IsZero() bool {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			if A.At(i, j).Cmp(zero) != 0 {
				return false
			}
		}
	}
	return true
}

// MatMul computes the matrix product of A and B.
func (A *Mat) MatMul(B *Mat) *Mat {
	C := Zeros(A.rows, B.cols)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < B.cols; j++ {
			for k := 0; k < A.cols; k++ {
				// C[i,j] += A[i,k] * B[k,j], in-place
				c := C.At(i, j)
				c.Add(c, new(big.Rat).Mul(A.At(i, k), B.At(k, j)))
			}
		}
	}
	return C
}

func (A *Mat) MatMulVec(u *Vec) *Vec {
	v := ZerosVec(A.rows)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			vj := v.At(j)
			vj.Add(vj, new(big.Rat).Mul(A.At(i, j), u.At(i)))
		}
	}
	return v
}

// LUFact computes the LU factorization of the (square) matrix A
//
// Finds L, U such that A = LU and L is lower triangular and U is upper
// triangular.
//
// Panics (divide by zero) if LU factorization does not exist.
//
// Source: [FNC: https://tobydriscoll.net/fnc-julia/linsys/lu.html]
func (A *Mat) LUFact() (L, U *Mat) {
	if A.rows != A.cols {
		panic("LUFact: matrix is not square")
	}
	n := A.rows

	// L is nxn
	// 1 0 1
	// 0 1 0
	// 0 0 1
	L = Zeros(n, n)
	for i := 0; i < n; i++ {
		L.Set(i, i, one)
	}

	U = Zeros(n, n)
	Ak := A.Copy()

	for k := 0; k < n-1; k++ {
		// U[k,:] = Ak[k,:]
		U.SetRow(k, Ak.Row(k))

		// L[:,k] = Ak[:,k] / U[k,k]
		L.SetCol(k, Ak.Col(k))
		for i := 0; i < n; i++ {
			l := L.At(i, k)
			L.At(i, k).Quo(l, U.At(k, k))
		}

		// Ak = Ak - L[:,k] * U[k,:]' (outer product)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// in-place subtraction
				a := Ak.At(i, j)
				a.Sub(a, new(big.Rat).Mul(L.At(i, k), U.At(k, j)))
			}
		}
	}

	U.Set(n-1, n-1, Ak.At(n-1, n-1))
	return L, U
}

// PLUFact computes the PLU factorization of the (square) matrix A
// See https://tobydriscoll.net/fnc-julia/linsys/pivoting.html
func (A *Mat) PLUFact() (L, U *Mat, p []int) {
	if A.rows != A.cols {
		panic("PLUFact: matrix is not square")
	}
	n := A.rows

	L = Zeros(n, n)
	U = Zeros(n, n)
	p = make([]int, n)
	Ak := A.Copy()

	// Selects row number in column k with maximum absolute value
	selectPivot := func(k int) int {
		pivot := Ak.Col(k).Argmax(func(r *big.Rat) *big.Rat { return new(big.Rat).Abs(r) })
		return pivot
	}

	for k := 0; k < n-1; k++ {
		// p[k] = argmax(abs(Ak[:,k]))
		p[k] = selectPivot(k)

		// U[k,:] = Ak[p[k],:]
		for j := 0; j < n; j++ {
			U.Set(k, j, Ak.At(p[k], j))
		}

		// L[:,k] = Ak[:,k] / U[k,k]
		for i := 0; i < n; i++ {
			l := new(big.Rat).Quo(Ak.At(i, k), U.At(k, k))
			L.Set(i, k, l)
		}

		// Ak = Ak - L[:,k] * U[k,:]' (outer product)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				a := Ak.At(i, j)
				a.Sub(a, new(big.Rat).Mul(L.At(i, k), U.At(k, j)))
			}
		}
	}
	p[n-1] = selectPivot(n - 1)
	// U[n,n] = Ak[p[n],n]
	U.Set(n-1, n-1, Ak.At(p[n-1], n-1))

	// L[:,n] = Ak[:,n] / U[n,n]
	for i := 0; i < n; i++ {
		l := new(big.Rat).Quo(Ak.At(i, n-1), U.At(n-1, n-1))
		L.Set(i, n-1, l)
	}
	return L.Pivot(p), U, p
}

func (u *Vec) Argmax(fn func(r *big.Rat) *big.Rat) int {
	max := fn(u.At(0))
	maxarg := 0
	for i := 1; i < u.rows; i++ {
		this := fn(u.At(i))
		if this.Cmp(max) > 0 {
			max = this
			maxarg = i
		}
	}
	return maxarg
}

func (u *Vec) Abs() *Vec {
	v := ZerosVec(u.rows)
	for i := 0; i < u.rows; i++ {
		v.Set(i, new(big.Rat).Abs(u.At(i)))
	}
	return v
}

// RowPermutation returns a new matrix with rows permuted according to p,
// as returned by PLUFact.
//
// Equivalent to A[p,:]
func (A *Mat) Pivot(p []int) *Mat {
	C := Zeros(A.rows, A.cols)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			C.Set(i, j, A.At(p[i], j))
		}
	}
	return C
}

func (u *Vec) Pivot(p []int) *Vec {
	v := ZerosVec(u.rows)
	for i := 0; i < u.rows; i++ {
		v.Set(i, u.At(p[i]))
	}
	return v
}

// Forwardsub solves a lower-triangular system Lx = b
// using forward substitution.
//
// See https://tobydriscoll.net/fnc-julia/linsys/linear-systems.html
func (L *Mat) Forwardsub(b *Vec) *Vec {
	n := L.rows
	// x = zeros(n)
	x := ZerosVec(n)
	// x[1] = b[1] / L[1,1]
	x.Set(0, new(big.Rat).Quo(b.At(0), L.At(0, 0)))
	// for i in 2:n
	for i := 1; i < n; i++ {
		// s = sum(L[i,j] * x[j] for j in 1:i-1)
		s := new(big.Rat)
		for j := 0; j < i; j++ {
			// s += L[i,j] * x[j]
			s.Add(s, new(big.Rat).Mul(L.At(i, j), x.At(j)))
		}
		// x[i] = (b[i] - s) / L[i,i]
		x.Set(i, new(big.Rat).Quo(new(big.Rat).Sub(b.At(i), s), L.At(i, i)))
	}
	return x
}

func (U *Mat) Backsub(b *Vec) *Vec {
	n := U.rows
	x := ZerosVec(n)
	// x[n] = b[n] / U[n,n]
	x.Set(n-1, new(big.Rat).Quo(b.At(n-1), U.At(n-1, n-1)))

	// for i in n-1:-1:1
	for i := n - 2; i >= 0; i-- {
		// s = sum(U[i,j] * x[j] for j in i+1:n)
		s := new(big.Rat)
		for j := i + 1; j < n; j++ {
			s.Add(s, new(big.Rat).Mul(U.At(i, j), x.At(j)))
		}
		// x[i] = (b[i] - s) / U[i,i]
		x.Set(i, new(big.Rat).Quo(new(big.Rat).Sub(b.At(i), s), U.At(i, i)))
	}
	return x
}

func (A *Mat) Backslash(b *Vec) *Vec {
	L, U, p := A.PLUFact()
	z := L.Forwardsub(b.Pivot(p))
	x := U.Backsub(z)
	return x
}

func (u *Vec) Equal(v *Vec) bool {
	for i := 0; i < u.rows; i++ {
		if u.At(i).Cmp(v.At(i)) != 0 {
			return false
		}
	}
	return true
}
