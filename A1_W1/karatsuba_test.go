package A1_W1

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionString(t *testing.T) {
	assert := assert.New(t)
	str1, str2 := partitionString("siema", 3)
	assert.Equal("sie", str1)
	assert.Equal("ma", str2)
}

func TestKaratsuba(t *testing.T) {
	assert := assert.New(t)
	res := Karatsuba(*big.NewInt(12), *big.NewInt(12))
	assert.Equal(0, big.NewInt(144).Cmp(&res))
}
