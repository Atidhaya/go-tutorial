package exercise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargestNumber(t *testing.T) {

	t.Run("Happy Path", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		result, err := KthLargestNumber(numbers, 1)
		assert.Equal(t, result, 6)
		assert.Nil(t, err)
	})

	t.Run("Happy Path 2", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		result, err := KthLargestNumber(numbers, 4)
		assert.Equal(t, result, 3)
		assert.Nil(t, err)
	})

	// //you can do subtest
	t.Run("Bad Path", func(t *testing.T) {
		t.Run("Negative Should Return Error", func(t *testing.T) {
			numbers := []int{1, 2, 3, 4, 5, 6}
			_, err := KthLargestNumber(numbers, -1)
			assert.Error(t, err)
		})

		t.Run("K exceed array length Should Return Error", func(t *testing.T) {
			numbers := []int{1, 2, 3, 4, 5, 6}
			_, err := KthLargestNumber(numbers, 10)
			assert.Error(t, err)
		})
	})

}
