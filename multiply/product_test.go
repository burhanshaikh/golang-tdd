package main

import "testing"

func TestProduct(t *testing.T) {

	t.Run("Printing the product of two integers", func(t *testing.T) {
		got := product(7, 9)
		want := 63

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})

}
