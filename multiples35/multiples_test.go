package multiples35

import "testing"

func TestMultiples(t *testing.T) {
	t.Run("returns the sum of all multiples of 3 and 5 below a threshold", func(t *testing.T) {
		t.Run("returns all the multiples of 3 and 5 below 16", func(t *testing.T) {
			got := Multiples35(16)
			want := 60
			asserEqual(t, got, want)
		})
		t.Run("returns all the multiples of 3 and 5 below 20", func(t *testing.T) {
			got := Multiples35(20)
			want := 78
			asserEqual(t, got, want)
		})
		t.Run("returns all the multiples of 3 and 5 below 20", func(t *testing.T) {
			got := Multiples35(10)
			want := 23
			asserEqual(t, got, want)
		})
	})
}

func asserEqual(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("wanted %d but got %d", want, got)
	}
}
