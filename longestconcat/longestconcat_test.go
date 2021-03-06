package longestconcat_test

import (
	"github.com/mikejeuga/katas/longestconcat"
	"reflect"
	"testing"
)

func TestLongestConcat(t *testing.T) {
	t.Run("returns the longest string of the concatenated slice given k=2", func(t *testing.T) {
		stringsInput := []string{"Alfie", "Dog", "Cat", "Fish", "Dragon"}
		want := "FishDragon"

		got := longestconcat.LongestConcat(stringsInput, 2)

		asserEqual(t, got, want)

	})

	t.Run("returns the longest string of the concatenated slice given k=3", func(t *testing.T) {
		stringsInput := []string{"Alfie", "Dog", "Cat", "Animal", "Bird", "Fish", "Dragon"}
		want := "AnimalBirdFish"

		got := longestconcat.LongestConcat(stringsInput, 3)

		asserEqual(t, got, want)

	})

	t.Run("Returns nothing if k is larger than the length of the slice", func(t *testing.T) {
		stringsInput := []string{"Alfie", "Dog", "Cat", "Animal", "Bird", "Fish", "Dragon"}
		want := ""

		got := longestconcat.LongestConcat(stringsInput, 10)

		asserEqual(t, got, want)
	})

}

func TestConcat(t *testing.T) {
	t.Run("returns a slice of the 2 concatenated strings from the input slice", func(t *testing.T) {
		stringsInput := []string{"Alfie", "Dog", "Cat", "Fish", "Dragon"}
		want := []string{"AlfieDog", "DogCat", "CatFish", "FishDragon"}

		got := longestconcat.Concat(stringsInput, 2)

		asserEqual(t, got, want)
	})

	t.Run("returns a slice of the 3 concatenated strings from the input slice", func(t *testing.T) {
		stringsInput := []string{"Alfie", "Dog", "Cat", "Animal", "Bird", "Fish", "Dragon"}
		want := []string{"AlfieDogCat", "DogCatAnimal", "CatAnimalBird", "AnimalBirdFish", "BirdFishDragon"}

		got := longestconcat.Concat(stringsInput, 3)

		asserEqual(t, got, want)
	})
}

func asserEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
