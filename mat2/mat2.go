package mat2

import "github.com/davecgh/go-spew/spew"

type Mat2 struct {
	A, B, C, D float64
}

// Solve finds x such that Ax + b = 0
func (m Mat2) Solve(b Vec2) Vec2 {
	// A B  @ b1
	// C D    b2

	// 1 0 @ A B  @ b1
	// 0 1   C D    b2

	// [1  B/A]  @ b1/A
	// [C   D ]     b2

	// [1       B/A  ]  @     b1/A
	// [0   D - C*B/A]     b2 - C*b1/A

	// [1       B/A  ]  @     b1/A
	// [0   1        ]     (b2 - C*b1/A)/(D - C*B/A)

	// (b2 - C*b1/A)/(D - C*B/A) = (b2 - C*b1/A)*(D + C*B/A)/[(D - C*B/A)(D + C*B/A)]
	// = b2*D + b2*C*B/A - C*b1/A*D - C*b1/A*C*B/A / [D^2  - C^2 * B^2 / A^2]
	// = b2*D + b2*C*B/A - C*b1/A*D - A^2*C*b1/A*C*B/A / [A^2 * D^2  - C^2 * B^2]
	// = D*b2 + B*C/A*b2 - C*D/A*b1 - B*C^2*b1 / [A^2 * D^2  - C^2 * B^2]

	// [1       0    ]  @     b1/A - B/A * (b2 - C*b1/A)/(D - C*B/A)
	// [0           1]     (b2 - C*b1/A)/(D - C*B/A)

	disc := m.A*m.D - m.B*m.C
	if disc == 0.0 {
		spew.Dump(m, b)
		panic("unsolvable")
	}

	// [ A B ] [ b1 ]
	// [ C D ] [ b2 ]
	// [ 1 2 ] [ 1 ]
	// [ 3 4 ] [ 2 ]

	m.B /= m.A
	b.B1 /= m.A
	m.A = 1 // 1 / A

	// [ 1 B ] [ b1 ]
	// [ C D ] [ b2 ]
	// [ 1 2 ] [ 1 ]
	// [ 3 4 ] [ 2 ]

	m.D -= m.B * m.C
	b.B2 -= b.B1 * m.C
	m.C = 0 // 2 - 1 * C

	// [ 1 B ] [ b1 ]
	// [ 0 D ] [ b2 ]
	// [ 1  2 ] [  1 ]
	// [ 0 -2 ] [ -1 ]

	m.C /= m.D // 2 / D
	b.B2 /= m.D
	m.D = 1

	// [ 1 B ] [ b1 ]
	// [ 0 1 ] [ b2 ]
	// [ 1 2 ] [  1  ]
	// [ 0 1 ] [ 1/2 ]

	b.B1 -= b.B2 * m.B
	m.B = 1

	// [ 1 0 ] [ b1 ]
	// [ 0 1 ] [ b2 ]
	// [ 1 0 ] [  0  ]
	// [ 0 1 ] [ 1/2 ]

	return b
}
