package A1_W1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHalveAr(t *testing.T) {
	assert := assert.New(t)
	ar1, ar2 := halveAr([]int{1, 2, 3, 4, 5})
	assert.Equal(2, len(ar1))
	assert.Equal(3, len(ar2))
	ar1, ar2 = halveAr([]int{1, 2, 3, 4})
	assert.Equal(2, len(ar1))
	assert.Equal(2, len(ar2))
}

func TestMergesort(t *testing.T) {
	assert := assert.New(t)
	res := MergeSort([]int{3, 1, 6, 4})
	assert.Equal(4, len(res))
	assert.Equal(1, res[0])
	assert.Equal(3, res[1])
	assert.Equal(4, res[2])
	assert.Equal(6, res[3])
}
