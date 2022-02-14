package stepsprimes_test

import (
	"github.com/mikejeuga/katas/stepsprimes"
	"reflect"
	"testing"
)

func TestStepsPrime(t *testing.T) {
	t.Run("steps in primes of 2", func(t *testing.T) {
		steps, start, end := 2, 100, 110

		got := stepsprimes.StepsPrimes(steps, start, end)
		want := []int{101, 103}

		assertEqual(t,got, want)
	})
	t.Run("steps in primes of 6", func(t *testing.T) {
		steps, start, end := 6, 100, 110

		got := stepsprimes.StepsPrimes(steps, start, end)
		want := []int{101, 107}

		assertEqual(t,got, want)
	})
	t.Run("steps in primes of 4", func(t *testing.T) {
		steps, start, end := 4, 100, 110

		got := stepsprimes.StepsPrimes(steps, start, end)
		want := []int{103, 107}

		assertEqual(t,got, want)
	})
}


func assertEqual(t *testing.T, got, want []int)  {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %q but got %q", want, got)
	}
}