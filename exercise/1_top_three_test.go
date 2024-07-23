package exercise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopThree(t *testing.T) {

	t.Run("Happy Path", func(t *testing.T) {
		expected := TopThreeWords{
			First:  "the",
			Second: "a",
			Third:  "of",
		}
		result, err := CalculateTopThree("wordcount.txt")
		assert.Equal(t, result, expected)
		assert.Nil(t, err)
	})

	t.Run("Bad Path", func(t *testing.T) {
		t.Run("file does not exist", func(t *testing.T) {
			_, err := CalculateTopThree("bar.txt")
			assert.Error(t, err)
			assert.Equal(t, "file bar.txt does not exist", err.Error(), "Error message does not match")
		})

		t.Run("file have less than 3 words", func(t *testing.T) {
			_, err := CalculateTopThree("two_word.txt")
			assert.Error(t, err)
			assert.Equal(t, "not enough words", err.Error(), "Error message does not match")
		})
	})

}
