package mathutil

import "testing"
import "github.com/stretchr/testify/assert"

func TestMulMod(t *testing.T) {

	assert.Equal(t,
		int64(16),
		MulMod(
			int64(4),
			int64(4),
			int64(17)), "")

	assert.Equal(t,
		int64(3),
		MulMod(
			int64(5),
			int64(4),
			int64(17)), "")

	assert.Equal(t,
		int64(6474921690063257),
		MulMod(
			int64(101234652348789987),
			int64(65746311545646431),
			int64(10005412336548794)),
		"")
}
