package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("Test plus", func(t *testing.T) {
		sum := 1 + 1
		if 2 != sum {
			t.Fail()
		}
	})
	t.Run("Test subtract", func(t *testing.T) {
		substract := 2 - 1
		if 1 != substract {
			t.Fail()
		}
	})
}
